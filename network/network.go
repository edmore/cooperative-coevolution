/*
Package network implements an artificial neural network
*/

package network

import (
	"github.com/edmore/esp/neuron"
	"github.com/edmore/esp/population"
)

type Network interface {
	activate()
}

type FeedForward struct {
	Activation  []float64
	HiddenUnits []*neuron.Neuron
	NumInputs   int
	NumOutputs  int
	Bias        float64
	Trials      int
	Fitness     int
	Parent1     int
	Parent2     int
	Id          int
}

func (f *FeedForward) activate() {
}

var counter int = 0

// FeedForward Network constructor
func NewFeedForward(in int, hid int, out int) *FeedForward {
	counter++
	return &FeedForward{
		Id:          counter,
		Activation:  make([]float64, hid),
		HiddenUnits: make([]*neuron.Neuron, hid),
		NumInputs:   in,
		NumOutputs:  out,
		Bias:        0.0,
		Trials:      0,
		Fitness:     0.0,
		Parent1:     -1,
		Parent2:     -1}
}

// Create the hidden units by randomly selecting them
func (f *FeedForward) Create(pops []*population.Population) {
	for i := 0; i < len(pops); i++ {
		f.HiddenUnits[i] = pops[i].SelectNeuron()
	}
}
