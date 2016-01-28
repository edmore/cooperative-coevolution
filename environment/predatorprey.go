/*
 Predator Prey Task
*/

package environment

import (
	//"fmt"
	//	"math"
	"github.com/edmore/esp/network"
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
	State *State
	World *Gridworld
}

// PredatorPrey Environment constructor
func NewPredatorPrey() *PredatorPrey {
	return &PredatorPrey{
		Name:  "Predator Prey Task",
		State: new(State),
		World: new(Gridworld)}
}

// Re-initialize the environment
func (p *PredatorPrey) Reset(n int) {
	// initialize world
	p.World.Length = 100
	p.World.Height = 100

	// initialise prey
	p.State.PreyX = 0
	p.State.PreyY = 99

	// initialize predators
	for i := 0; i < n; i++ {
		p.State.PredatorX = append(p.State.PredatorX, i*2)
		p.State.PredatorY = append(p.State.PredatorY, 0)
	}
}

func (p *PredatorPrey) PerformPredatorAction(predator network.Network, action []float64) {

}

func (p *PredatorPrey) PerformPreyAction(state *State) {
}

// Get the current state variables
func (p *PredatorPrey) GetState() *State {
	return p.State
}

func (p *PredatorPrey) GetWorld() *Gridworld {
	return p.World
}

// Prey Caught
func (p *PredatorPrey) Caught() bool {
	return false
}

// Prey Surrounded
func (p *PredatorPrey) Surrounded() bool {
	return false
}
