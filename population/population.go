/*
Package population implements a population of neurons
*/

package population

import (
	"github.com/edmore/esp/neuron"
	"math/rand"
)

type Population struct {
	Id          int
	Neurons     []*neuron.Neuron
	Individuals int
	Evolvable   bool
	Numbreed    int
}

var counter int = 0

// Population constructor
func NewPopulation(size int) *Population {
	counter++
	return &Population{
		Id:          counter,
		Individuals: size,
		Evolvable:   true,
		Numbreed:    size / 4,
		Neurons:     make([]*neuron.Neuron, size)}
}

// Create the neurons, put them in the (sub)population and initialize their weights
func (p *Population) Create() {
	if p.Evolvable {
		for i := 0; i < p.Individuals; i++ {
			p.Neurons[i] = neuron.NewNeuron(p.Individuals)
			p.Neurons[i].Create()
		}
	}
}

// Select a neuron at random
func (p *Population) SelectNeuron() *neuron.Neuron {
	index := rand.Int() % p.Individuals
	return p.Neurons[index]
}
