/*
Package network implements a recurrent artificial neural network
*/

package network

import (
	"github.com/edmore/esp/activation/sigmoid"
	"github.com/edmore/esp/neuron"
	"github.com/edmore/esp/population"
)

type Recurrent struct {
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
func NewRecurrent(in int, hid int, out int, bias bool) *Recurrent {
	counter++
	genesize := in + out
	if bias == true {
		genesize++
	}

	return &Recurrent{
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
func (r *Recurrent) Activate(input []float64, output []float64) []float64 {
	// input layer -> hidden layer
	for key, neuron := range r.HiddenUnits {
		if !neuron.Lesioned {
			for i := 0; i < len(input); i++ {
				r.Activation[key] = r.Activation[key] + (neuron.Weight[i] * input[i])
			}
			r.Activation[key] = sigmoid.Logistic(1.0, r.Activation[key])
		}
	}
	// hidden layer -> output layer
	for i := 0; i < r.NumOutputs; i++ {
		for key, neuron := range r.HiddenUnits {
			output[i] = output[i] + (r.Activation[key] * neuron.Weight[len(input)+i])
		}
		output[i] = sigmoid.Logistic(1.0, output[i])
	}
	return output
}

// Return the hidden units
func (r *Recurrent) GetHiddenUnits() []*neuron.Neuron {
	return r.HiddenUnits
}

// Create the hidden units by randomly selecting them
func (r *Recurrent) Create(pops []*population.Population) {
	for i := 0; i < len(pops); i++ {
		r.HiddenUnits[i] = pops[i].SelectNeuron()
	}
}

// Return the total number of inputs
func (r *Recurrent) GetTotalInputs() int {
	if r.Bias == true {
		return r.NumInputs + 1
	} else {
		return r.NumInputs
	}
}

// Return the total number of outputs
func (r *Recurrent) GetTotalOutputs() int {
	return r.NumOutputs
}

// Return true if network has bias
func (r *Recurrent) HasBias() bool {
	return r.Bias
}

// Set the fitness for a network
func (r *Recurrent) SetFitness(fitness int) {
	r.Fitness = fitness
}

// Get the fitness for a network
func (r *Recurrent) GetFitness() int {
	return r.Fitness
}

// Increment the cumulative fitness and trials for the network neurons
func (r *Recurrent) SetNeuronFitness() {
	for _, neuron := range r.HiddenUnits {
		neuron.SetFitness(r.Fitness)
		neuron.Trials++
	}
}

// Tag best neurons
func (r *Recurrent) Tag() {
	for _, neuron := range r.HiddenUnits {
		neuron.Tag = true
	}
}

// Reset the network activation
func (r *Recurrent) ResetActivation() {
	r.Activation = make([]float64, len(r.GetHiddenUnits()))
}

// Reset the network fitness and trials
func (r *Recurrent) ResetFitness() {
	r.Fitness, r.Trials = 0, 0
}
