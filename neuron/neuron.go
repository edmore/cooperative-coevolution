/*
Package neuron implements an Artificial Neural Network neuron.
*/

package neuron

import (
	"math/rand"
	"github.com/edmore/esp/random"
)

type Neuron struct {
	weight   []float64
	lesioned bool
	trials   int
	fitness  float64
	tag      bool
	parent1  int
	parent2  int
	name     string
	id       int
}

var counter int = 0

// Neuron constructor
func NewNeuron(size int) *Neuron {
	counter++
	return &Neuron{
		id:      counter,
		weight:  make([]float64, size),
		name:    "basic neuron",
		parent1: -1,
		parent2: -1}
}

// Create a new set of random weights
func (n *Neuron) Create() {
	for i := 0; i < len(n.weight); i++ {
		n.weight[i] = (rand.Float64() * 12.0) - 6.0
	}
}

// Perturb the weights of a Neuron.
// Used to search in a neighborhood around some Neuron (best).
func (n *Neuron) Perturb() {
	coefficient := 0.3
	for i := 0; i < len(n.weight); i++ {
		n.weight[i] = n.weight[i] + random.Cauchy(coefficient)
	}
	// reset fitness and trials
	n.fitness, n.trials = 0, 0
}
