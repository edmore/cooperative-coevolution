/*
 Package cartpole implements the Double Pole balancing / Inverted Pendulum Task.
 Currently - with full state information (markov)
*/

package cartpole

import (
	"github.com/edmore/esp/rungekutta/rk4"

	"math"
)

var (
	ForceMag float64 = 10.0 // Fixed Force magnitude i.e. (+ / -) 10.0
	Tau      float64 = 0.02 //seconds between state updates (the time step)
)

const (
	RadToDeg         = 180 / math.Pi
	DegToRad         = math.Pi / 180
	Gravity  float64 = 9.81
)

type State struct {
	X         float64 // position of the cart
	XDot      float64 // velocity of the cart
	Theta1    float64 // angle of the 1st pole
	ThetaDot1 float64 // angular velocity of the 1st pole
	Theta2    float64 // angle of the 2nd pole
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

// Runge-Kutta Step - approximate state variables at time dt
func step(action int, c *Cartpole) {
	dt := 0.01 // step size
	// Equations for cart position and pole angles
	eq1 := func(x, y float64) float64 { return state.XDot }
	eq2 := func(x, y float64) float64 { return state.ThetaDot1 }
	eq3 := func(x, y float64) float64 { return state.ThetaDot2 }
	if action > 0 {
		//	F := ForceMag
	} else {
		//	F := -ForceMag
	}

	// update position of cart
	pt := rk4.NewPoint(0, state.X)
	state.X = pt.Solve(dt, eq1, Tau)

	// update angles of the poles
	pt = rk4.NewPoint(0, state.Theta1)
	state.Theta1 = pt.Solve(dt, eq2, Tau)

	pt = rk4.NewPoint(0, state.Theta2)
	state.Theta2 = pt.Solve(dt, eq3, Tau)
}

// Get the current state variables
func (c *Cartpole) GetState() *State {
	return state
}
