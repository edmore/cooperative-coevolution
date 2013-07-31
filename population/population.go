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

func NewPopulation(size int) *Population {
	counter++
	return &Population{
		id:          counter,
		individuals: size,
		evolvable:   true,
		numbreed:    size / 4,
		neurons:     make([]*neuron.Neuron, size)}
}

func (p *Population) Create() {
	for i := 0; i < p.individuals; i++ {
		p.neurons[i] = neuron.NewNeuron(p.individuals)
		p.neurons[i].Create()
	}
}
