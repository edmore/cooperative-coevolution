/*
Package network implements an artificial neural network
*/

package network

import (
	"github.com/edmore/esp/neuron"
	"github.com/edmore/esp/population"
)

type FeedForward struct {
	Id          int
	Activation  []float64
	HiddenUnits []*neuron.Neuron
	NumInputs   int
	NumOutputs  int
	Bias        bool
	Trials      int
	Fitness     float64
	Parent1     int
	Parent2     int
	Name        string
	GeneSize    int
}

var counter int = 0

// FeedForward Network constructor
func NewFeedForward(in int, hid int, out int, bias bool) *FeedForward {
	counter++
	genesize := in + out
	if bias == true {
		genesize++
	}

	return &FeedForward{
		Id:          counter,
		Activation:  make([]float64, hid),
		HiddenUnits: make([]*neuron.Neuron, hid),
		NumInputs:   in,
		NumOutputs:  out,
		Bias:        bias,
		Parent1:     -1,
		Parent2:     -1,
		Name:        "Feed Forward",
		GeneSize:    genesize}
}

// Activate
func (f *FeedForward) Activate(inputs []float64) float64 {
	return 1
}

// Return the hidden units
func (f *FeedForward) GetHiddenUnits() []*neuron.Neuron {
	return f.HiddenUnits
}

// Create the hidden units by randomly selecting them
func (f *FeedForward) Create(pops []*population.Population) {
	for i := 0; i < len(pops); i++ {
		f.HiddenUnits[i] = pops[i].SelectNeuron()
	}
}

// Return the total number of inputs
func (f *FeedForward) GetTotalInputs() int {
	if f.Bias == true {
		return f.NumInputs + 1
	} else {
		return f.NumInputs
	}
}

// Return true if network has bias
func (f *FeedForward) HasBias() bool {
	return f.Bias
}
