/*
Package population implements a population of neurons
*/

package population

import (
	"github.com/edmore/esp/neuron"
)

type Population struct {
	id          int
	neurons     []*neuron.Neuron
	individuals int
	evolvable   bool
	numbreed    int
}

var counter int = 0

// Population constructor
func NewPopulation(size int) *Population {
	counter++
	return &Population{
		id:          counter,
		individuals: size,
		evolvable:   true,
		numbreed:    size / 4,
		neurons:     make([]*neuron.Neuron, size)}
}

// Create the neurons, put them in the (sub)population and initialize their weights
func (p *Population) Create() {
	if p.evolvable {
		for i := 0; i < p.individuals; i++ {
			p.neurons[i] = neuron.NewNeuron(p.individuals)
			p.neurons[i].Create()
		}
	}
}
