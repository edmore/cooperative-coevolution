/*
 Predator Prey Task
*/

package environment

import (
	//"fmt"
	"math"
	"math/rand"
)

type Gridworld struct {
	Length int
	Height int
}

type State struct {
	PredatorX []int // x position(s) of the predator(s)
	PredatorY []int // y position(s) of the predator(s)
	PreyX     int   // x position of the prey
	PreyY     int   // y position of the prey
	Caught    bool
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

	// initialise prey randomly
	p.State.PreyX = rand.Intn(100)
	p.State.PreyY = rand.Intn(100)

	// initialize predators
	for i := 0; i < n; i++ {
		p.State.PredatorX = append(p.State.PredatorX, i*2)
		p.State.PredatorY = append(p.State.PredatorY, 0)
	}
	p.State.Caught = false
}

func (p *PredatorPrey) PerformPredatorAction(position int, action []float64) {
	// the action vector contains real numbers that map to coordinates
	// the position of the max number maps to the direction the predator should go
	predAction := getMaxPosition(action)

	if predAction == 0 {
		p.State.PredatorY[position]++ // N
	} else if predAction == 1 {
		p.State.PredatorX[position]++ // E
	} else if predAction == 2 {
		p.State.PredatorY[position]-- // S
	} else if predAction == 3 {
		p.State.PredatorX[position]-- // W
	}
	// else Stay

	// the toroidal wrap-around
	if p.State.PredatorX[position] > p.World.Length {
		p.State.PredatorX[position] = p.State.PredatorX[position] - p.World.Length
	}
	if p.State.PredatorY[position] > p.World.Height {
		p.State.PredatorY[position] = p.State.PredatorY[position] - p.World.Height
	}
	if p.State.PredatorX[position] < 0 {
		p.State.PredatorX[position] = p.State.PredatorX[position] + p.World.Length
	}
	if p.State.PredatorY[position] < 0 {
		p.State.PredatorY[position] = p.State.PredatorY[position] + p.World.Height
	}

	// Check if prey has been caught

	//	fmt.Println(p.State)
	if (p.State.PredatorX[position] == p.State.PreyX) && (p.State.PredatorY[position] == p.State.PreyY) {
		p.State.Caught = true
		//		fmt.Println("Yay")
	}
}

func getMaxPosition(action []float64) int {
	max := action[0]
	result := 0

	for i := 0; i < len(action); i++ {
		if action[i] > max {
			max = action[i]
			result = i
		}
	}
	return result
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

	// the toroidal wrap-around
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
	return p.State.Caught
}

// Prey Surrounded
func (p *PredatorPrey) Surrounded() bool {
	return false
}
