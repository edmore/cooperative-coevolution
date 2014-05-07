/*
Package population implements a population of neurons
*/

package population

import (
	"github.com/edmore/esp/neuron"

	"math/rand"
	"sort"
)

type Neurons []*neuron.Neuron

// for sorting neurons by average fitness
func (s Neurons) Len() int      { return len(s) }
func (s Neurons) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByAvgFitness implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Neurons value.
type ByAvgFitness struct{ Neurons }

func (s ByAvgFitness) Less(i, j int) bool {
	// sort in descending order - largest first
	return (s.Neurons[i].Fitness / s.Neurons[i].Trials) > (s.Neurons[j].Fitness / s.Neurons[j].Trials)
}

type Population struct {
	Id             int
	Individuals    Neurons
	NumIndividuals int
	Evolvable      bool
	Numbreed       int // the number of neurons to breed - top quartile
	GeneSize       int
}

var counter int = 0

// Population constructor
func NewPopulation(size int, genesize int) *Population {
	counter++
	return &Population{
		Id:             counter,
		NumIndividuals: size,
		Evolvable:      true,
		Numbreed:       size / 4,
		Individuals:    make(Neurons, size),
		GeneSize:       genesize}
}

// Create the neurons, put them in the (sub)population and initialize their weights
func (p *Population) Create() {
	if p.Evolvable {
		for i := 0; i < p.NumIndividuals; i++ {
			p.Individuals[i] = neuron.NewNeuron(p.GeneSize)
			p.Individuals[i].Create()
		}
	}
}

// Select a neuron at random
func (p *Population) SelectNeuron() *neuron.Neuron {
	index := rand.Int() % p.NumIndividuals
	return p.Individuals[index]
}

// Sort neurons in population
func (p *Population) SortNeurons() {
	sort.Sort(ByAvgFitness{p.Individuals})
}
