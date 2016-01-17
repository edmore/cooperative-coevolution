/*
 Predator Prey Task
*/

package environment

import (
//"fmt"
//	"math"
)

const ()

var ()

type Gridworld struct {
	Length int
	Height int
}

type State struct {
	PredatorX []int // x position(s) of the predator(s)
	PredatorY []int // y position(s) of the predator(s)
	PreyX     int   // x position of the prey
	PreyY     int   // y position of the prey
}

type PredatorPrey struct {
	Name  string
	state *State
	world *Gridworld
}

// PredatorPrey Environment constructor
func NewPredatorPrey() *PredatorPrey {
	return &PredatorPrey{
		Name:  "Predator Prey Task",
		state: new(State),
		world: new(Gridworld)}
}

// Re-initialize the environment
func (p *PredatorPrey) Reset() {
	// you can set the position of the predators here
}

// Stores the desired action for the next step
func (p *PredatorPrey) PerformPredatorAction(action float64) {
	predatorStep(action, p)
}

// Step
func predatorStep(action float64, p *PredatorPrey) {
}

// Stores the desired action for the next step
func (p *PredatorPrey) PerformPreyAction(action float64) {
	preyStep(action, p)
}

// Step
func preyStep(action float64, p *PredatorPrey) {
}

// Get the current state variables
func (p *PredatorPrey) GetState() *State {
	return p.state
}

func (p *PredatorPrey) GetWorld() *Gridworld {
	return p.world
}

// Prey Caught
func (p *PredatorPrey) Caught() bool {
	return false
}

// Prey Surrounded
func (p *PredatorPrey) Surrrounded() bool {
	return false
}
