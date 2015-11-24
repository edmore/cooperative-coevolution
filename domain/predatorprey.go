/*
 Predator Prey Task
*/

package predatorprey

import (
	//"fmt"
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

type PredatorPrey struct {
	Name      string
	TrackSize float64
	Up        float64 // Pole-hinge Friction Coefficient
	Uc        float64 // Cart-track Friction Coefficient
	MassCart  float64
	MassPole1 float64
	MassPole2 float64
	Length1   float64 // actually half the pole's length
	Length2   float64 // actually half the pole's length
	state     *State
}

// Cartpole Environment constructor
func NewPredatorPrey(length2 float64) *PredatorPrey {
	return &PredatorPrey{
		Name:      "Predator Prey Task",
		TrackSize: 2.4,
		Up:        0.000002,
		Uc:        0.0005,
		MassCart:  1.0,
		MassPole1: 0.1,
		MassPole2: 0.01,
		Length1:   0.5,
		Length2:   length2,
		state:     new(State)}
}

// Re-initialize the environment
func (p *PredatorPrey) Reset() {

}

// Stores the desired action for the next Runge-Kutta step
func (p *PredatorPrey) PerformPredatorAction(action float64) {
	predatorStep(action, c)
}

// Runge-Kutta Step - approximate state variables at time Tau
func predatorStep(action float64, p *PredatorPrey) {
}

// Stores the desired action for the next Runge-Kutta step
func (p *PredatorPrey) PerformPreyAction(action float64) {
	preyStep(action, c)
}

// Runge-Kutta Step - approximate state variables at time Tau
func preyStep(action float64, p *PredatorPrey) {
}

// Get the current state variables
func (p *PredatorPrey) GetState() *State {
	return p.state
}

// Prey Caught
func (p *PredatorPrey) Caught() bool {
	return false
}

// Prey Caught
func (p *PredatorPrey) Surrrounded() bool {
	return false
}
