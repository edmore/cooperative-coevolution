/*
 Predator Prey Task
*/

package environment

import (
	//"fmt"
	"github.com/edmore/esp/network"
	"math"
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
	p.State.PreyX = 6
	p.State.PreyY = 0

	// initialize predators
	for i := 0; i < n; i++ {
		p.State.PredatorX = append(p.State.PredatorX, i*2)
		p.State.PredatorY = append(p.State.PredatorY, 0)
	}
}

func (p *PredatorPrey) PerformPredatorAction(predator network.Network, action []float64) {

}

func (p *PredatorPrey) PerformPreyAction(nearest int) {
	// calculate nearest predator offset with wrap-around
	distanceX := float64(p.State.PredatorX[nearest] - p.State.PreyX)
	if math.Abs(distanceX) > float64(p.World.Length/2) {
		temp := distanceX
		distanceX = float64(p.World.Length) - math.Abs(distanceX)
		if temp > 0 {
			distanceX = 0 - distanceX
		}
	}

	distanceY := float64(p.State.PredatorY[nearest] - p.State.PreyY)
	if math.Abs(distanceY) > float64(p.World.Height/2) {
		temp := distanceY
		distanceY = float64(p.World.Height) - math.Abs(distanceY)
		if temp > 0 {
			distanceY = 0 - distanceY
		}
	}

	// Move N,S,E,W or Stay
	if distanceY < 0 && (math.Abs(float64(distanceY)) >= math.Abs(float64(distanceX))) {
		// Move N
		p.State.PreyY++
	} else if distanceX < 0 && (math.Abs(float64(distanceX)) >= math.Abs(float64(distanceY))) {
		// Move E
		p.State.PreyX++
	} else if distanceY > 0 && (math.Abs(float64(distanceY)) >= math.Abs(float64(distanceX))) {
		// Move S
		p.State.PreyY--
	} else if distanceX > 0 && (math.Abs(float64(distanceX)) >= math.Abs(float64(distanceY))) {
		// Move W
		p.State.PreyX--
	}
	// Else Stay

	// the toroid wrap-around
	if p.State.PreyX > p.World.Length {
		p.State.PreyX = p.State.PreyX - p.World.Length
	}
	if p.State.PreyY > p.World.Height {
		p.State.PreyY = p.State.PreyY - p.World.Height
	}
	if p.State.PreyX < 0 {
		p.State.PreyX = p.State.PreyX + p.World.Length
	}
	if p.State.PreyY < 0 {
		p.State.PreyY = p.State.PreyY + p.World.Height
	}

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
