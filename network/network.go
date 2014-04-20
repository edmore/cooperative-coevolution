package network

import "github.com/edmore/esp/neuron"

type Network interface {
	Activate([]float64) []float64
	GetHiddenUnits() []*neuron.Neuron
	GetTotalInputs() int
	HasBias() bool
	SetFitness(int)
	GetFitness() int
}
