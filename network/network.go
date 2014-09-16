package network

import (
	"github.com/edmore/esp/neuron"
	"github.com/edmore/esp/population"
)

type Network interface {
	Activate([]float64, []float64) []float64
	Create([]*population.Population)
	GetHiddenUnits() []*neuron.Neuron
	GetTotalInputs() int
	GetTotalOutputs() int
	HasBias() bool
	SetFitness(int)
	GetFitness() int
	SetNeuronFitness()
	ResetActivation()
	ResetFitness()
	Tag()
}

var counter int = 0
