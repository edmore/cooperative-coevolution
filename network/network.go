package network

import (
	"github.com/edmore/cooperative-coevolution/neuron"
	"github.com/edmore/cooperative-coevolution/population"
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
	GetID() int
}

var counter int = 0
