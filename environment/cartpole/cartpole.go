/*
 Package cartpole implements the Double Pole balancing / Inverted Pendulum Task.
 Currently - with full state information (markov)
*/

package cartpole

import (
	"github.com/edmore/esp/rungekutta/rk4"
	"math"
)

const (
	RadToDeg         = 180 / math.Pi
	DegToRad         = math.Pi / 180
	Gravity  float64 = 9.81
)

var (
	ForceMag     float64 = 10.0 // Fixed Force magnitude i.e. (+ / -) 10.0
	Tau          float64 = 0.02 //seconds between state updates (the time step)
	FailureAngle float64 = 36.0 // failure angle in degrees
)

type State struct {
	X         float64 // position of the cart
	XDot      float64 // velocity of the cart
	Theta1    float64 // angle of the 1st pole
	Theta2    float64 // angle of the 2nd pole
	ThetaDot1 float64 // angular velocity of the 1st pole
	ThetaDot2 float64 // angular velocity of the 2nd pole
}

var state *State = new(State)

type Cartpole struct {
	Name      string
	TrackSize float64
	Up        float64 // Pole-hinge Friction Coefficient
	Uc        float64 // Cart-track Friction Coefficient
	MassCart  float64
	MassPole1 float64
	MassPole2 float64
	Length1   float64 // actually half the pole's length
	Length2   float64 // actually half the pole's length
}

// Cartpole Environment constructor
func NewCartpole() *Cartpole {
	return &Cartpole{
		Name:      "Double Pole Balancing Task",
		TrackSize: 2.4,
		Up:        0.000002,
		Uc:        0.0005,
		MassCart:  1.0,
		MassPole1: 0.1,
		MassPole2: 0.01,
		Length1:   0.5,
		Length2:   0.05}
}

// Re-initialize the environment
func (c *Cartpole) Reset() {
	state.X = 0.2            // position of cart
	state.Theta1 = DegToRad  // angle of the 1st pole - 1 degree
	state.Theta2 = -DegToRad // angle of the 2nd pole - minus 1 degree
}

// Stores the desired action for the next Runge-Kutta step
func (c *Cartpole) PerformAction(action int) *State {
	step(action, c)
	return state
}

// Runge-Kutta Step - approximate state variables at time Tau
func step(action int, c *Cartpole) {
	dt := 0.01 // step size
	var F float64
	if action > 0 {
		F = ForceMag
	} else {
		F = -ForceMag
	}

	sinTheta1 := math.Sin(state.Theta1)
	cosTheta1 := math.Cos(state.Theta1)
	gSinTheta1 := Gravity * sinTheta1

	sinTheta2 := math.Sin(state.Theta2)
	cosTheta2 := math.Cos(state.Theta2)
	gSinTheta2 := Gravity * sinTheta2

	temp1 := c.Up * state.Theta1 / c.Length1 * c.MassPole1
	temp2 := c.Up * state.Theta2 / c.Length2 * c.MassPole2
	fi1 := (c.Length1 * c.MassPole1 * math.Pow(state.Theta1, 2) * sinTheta1) +
		(0.75 * c.MassPole1 * cosTheta1 * (temp1 + gSinTheta1))
	fi2 := (c.Length2 * c.MassPole2 * math.Pow(state.Theta2, 2) * sinTheta2) +
		(0.75 * c.MassPole2 * cosTheta2 * (temp2 + gSinTheta2))
	mi1 := c.MassPole1 * (1 - (0.75 * math.Pow(cosTheta1, 2)))
	mi2 := c.MassPole2 * (1 - (0.75 * math.Pow(cosTheta2, 2)))

	xDotDot := (F + fi1 + fi2) / (mi1 + mi2 + c.MassCart)
	thetaDotDot1 := -0.75 * (xDotDot*cosTheta1 + gSinTheta1 + temp1) / c.Length1
	thetaDotDot2 := -0.75 * (xDotDot*cosTheta2 + gSinTheta2 + temp2) / c.Length2

	// Equations for cart position and pole angles
	eq1 := func(x, y float64) float64 { return state.XDot }
	eq2 := func(x, y float64) float64 { return state.ThetaDot1 }
	eq3 := func(x, y float64) float64 { return state.ThetaDot2 }
	// Equations of motion derivatives
	eq4 := func(x, y float64) float64 { return xDotDot }
	eq5 := func(x, y float64) float64 { return thetaDotDot1 }
	eq6 := func(x, y float64) float64 { return thetaDotDot2 }

	// update position of cart
	pt := rk4.NewPoint(0, state.X)
	state.X = pt.Solve(dt, eq1, Tau)

	// update angles of the poles
	pt = rk4.NewPoint(0, state.Theta1)
	state.Theta1 = pt.Solve(dt, eq2, Tau)
	pt = rk4.NewPoint(0, state.Theta2)
	state.Theta2 = pt.Solve(dt, eq3, Tau)

	// update velocity of cart
	pt = rk4.NewPoint(0, state.XDot)
	state.XDot = pt.Solve(dt, eq4, Tau)

	// update angular velocities of the poles
	pt = rk4.NewPoint(0, state.ThetaDot1)
	state.ThetaDot1 = pt.Solve(dt, eq5, Tau)
	pt = rk4.NewPoint(0, state.ThetaDot2)
	state.ThetaDot2 = pt.Solve(dt, eq6, Tau)
}

// Get the current state variables
func (c *Cartpole) GetState() *State {
	return state
}

// Cart within track bounds
func (c *Cartpole) WithinTrackBounds() bool {
	return (state.X > -c.TrackSize && state.X < c.TrackSize)
}

// Pole angles within acceptable bounds
func (c *Cartpole) WithinAngleBounds() bool {
	failure := FailureAngle * DegToRad // ~ 0.6283185 radians
	return (state.Theta1 > -failure && state.Theta1 < failure) &&
		(state.Theta2 > -failure && state.Theta2 < failure)
}
