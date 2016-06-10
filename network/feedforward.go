/*
Package network implements a feedforward artificial neural network
*/

package network

import (
	"github.com/edmore/cooperative-coevolution/activation/sigmoid"
	"github.com/edmore/cooperative-coevolution/neuron"
	"github.com/edmore/cooperative-coevolution/population"
)

type FeedForward struct {
	Id          int
	Activation  []float64
	HiddenUnits []*neuron.Neuron
	NumInputs   int
	NumOutputs  int
	Bias        bool
	Trials      int
	Fitness     int
	Parent1     int
	Parent2     int
	Name        string
	GeneSize    int
}

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
func (f *FeedForward) Activate(input []float64, output []float64) []float64 {
	// input layer -> hidden layer
	for key, neuron := range f.HiddenUnits {
		if !neuron.Lesioned {
			for i := 0; i < len(input); i++ {
				f.Activation[key] = f.Activation[key] + (neuron.Weight[i] * input[i])
			}
			f.Activation[key] = sigmoid.Logistic(1.0, f.Activation[key])
		}
	}
	// hidden layer -> output layer
	for i := 0; i < f.NumOutputs; i++ {
		for key, neuron := range f.HiddenUnits {
			output[i] = output[i] + (f.Activation[key] * neuron.Weight[len(input)+i])
		}
		output[i] = sigmoid.Logistic(1.0, output[i])
	}
	return output
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

// Return the total number of outputs
func (f *FeedForward) GetTotalOutputs() int {
	return f.NumOutputs
}

// Return true if network has bias
func (f *FeedForward) HasBias() bool {
	return f.Bias
}

// Set the fitness for a network
func (f *FeedForward) SetFitness(fitness int) {
	f.Fitness = fitness
}

// Get the fitness for a network
func (f *FeedForward) GetFitness() int {
	return f.Fitness
}

// Increment the cumulative fitness and trials for the network neurons
func (f *FeedForward) SetNeuronFitness() {
	for _, neuron := range f.HiddenUnits {
		neuron.SetFitness(f.Fitness)
		neuron.Trials++
	}
}

// Tag best neurons
func (f *FeedForward) Tag() {
	for _, neuron := range f.HiddenUnits {
		neuron.Tag = true
	}
}

// Reset the network activation
func (f *FeedForward) ResetActivation() {
	f.Activation = make([]float64, len(f.GetHiddenUnits()))
}

// Reset the network fitness and trials
func (f *FeedForward) ResetFitness() {
	f.Fitness, f.Trials = 0, 0
}
