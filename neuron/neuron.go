/*
Package neuron implements an Artificial Neural Network neuron.
*/

package neuron

import "math/rand"

type Neuron struct {
	weight   []float64
	lesioned bool
	trials   int
	fitness  float64
}

// Neuron constructor
func NewNeuron(size int) *Neuron {
	var n *Neuron = new(Neuron)
	n.weight = make([]float64, size)
	return n
}

// Create a new set of random weights
func (n *Neuron) Create() {
	for i := 0; i < len(n.weight); i++ {
		n.weight[i] = (rand.Float64() * 12.0) - 6.0
	}
}
