/*
Package neuron implements an Artificial Neural Network neuron.
*/

package neuron

import (
	"github.com/edmore/esp/random"
	"math/rand"
)

type Neuron struct {
	Weight   []float64
	Lesioned bool
	Trials   int
	Fitness  int
	Tag      bool
	Parent1  int
	Parent2  int
	Name     string
	Id       int
}

var counter int = 0

// Neuron constructor
func NewNeuron(size int) *Neuron {
	counter++
	return &Neuron{
		Id:      counter,
		Weight:  make([]float64, size),
		Name:    "basic neuron",
		Parent1: -1,
		Parent2: -1}
}

// Create a new set of random weights
func (n *Neuron) Create() {
	for i := 0; i < len(n.Weight); i++ {
		n.Weight[i] = (rand.Float64() * 12.0) - 6.0
	}
}

// Set the (cumulative) fitness for the neuron
func (n *Neuron) SetFitness(fitness int) {
	n.Fitness = n.Fitness + fitness
}

// Perturb the weights of a Neuron.
// Used to search in a neighborhood around some Neuron (best).
func (n *Neuron) Perturb() {
	coefficient := 0.3
	for i := 0; i < len(n.Weight); i++ {
		n.Weight[i] = n.Weight[i] + random.Cauchy(coefficient)
	}
	// reset fitness and trials
	n.Fitness, n.Trials = 0, 0
}

// Reset Fitness and Trials
func (n *Neuron) ResetFitness() {
	n.Fitness, n.Trials = 0, 0
}
