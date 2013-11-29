/*
 Package cartpole implements the Double Pole balancing / Inverted Pendulum Task.
 Currently - with full state information (markov)
*/

package cartpole

import "math"

var (
	ForceMag float64 = 10.0 // Force magnitude - (+ / -) 10.0
	Tau      float64 = 0.01 //seconds between state updates (the time step)
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
func (c *Cartpole) Reset() *State {
	state.X = 0.2
	state.Theta1 = DegToRad
	return state
}

// Stores the desired action for the next Runge-Kutta step
func (c *Cartpole) PerformAction(action int) *State {
	state := step(action)
	return state
}

// Runge-Kutta Step - approximate state variables at time dt
func step(action int) *State {
	if action == 0 {
		F := ForceMag * -1
	} else {
		F := ForceMag
	}
	return state
}
