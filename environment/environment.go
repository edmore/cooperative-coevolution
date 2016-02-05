package environment

import "github.com/edmore/esp/network"

type Environment interface {
	GetWorld() *Gridworld
	GetState() *State
	Caught() bool
	Surrounded() bool
	PerformPreyAction(int)
	PerformPredatorAction(network.Network, []float64)
	Reset(int)
}
