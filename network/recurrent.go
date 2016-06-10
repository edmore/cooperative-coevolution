/*
Package network implements a recurrent artificial neural network
*/

package network

import (
	"github.com/edmore/cooperative-coevolution/activation/sigmoid"
	"github.com/edmore/cooperative-coevolution/neuron"
	"github.com/edmore/cooperative-coevolution/population"
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

// Recurrent Network constructor
func NewRecurrent(in int, hid int, out int, bias bool) *Recurrent {
	counter++
	genesize := in + hid
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
		Name:        "Recurrent",
		GeneSize:    genesize}
}

// Activate
func (r *Recurrent) Activate(input []float64, output []float64) []float64 {
	delay := 2
	//tmp := make([]float64, len(input)+len(r.Activation))

	for d := 0; d < delay; d++ {
		tmp := input
		tmp = append(tmp, r.Activation...)

		// input layer -> hidden layer
		for key, neuron := range r.HiddenUnits {
			r.Activation[key] = 0.0
			if !neuron.Lesioned {
				for i := 0; i < len(input)+len(r.GetHiddenUnits()); i++ {
					r.Activation[key] = r.Activation[key] + (neuron.Weight[i] * tmp[i])
				}
				r.Activation[key] = sigmoid.Logistic(1.0, r.Activation[key])
			}
		}

		for i := 0; i < r.NumOutputs; i++ {
			output[i] = r.Activation[i]
		}
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
