/*
Package population implements a population of neurons
*/

package population

import (
	"github.com/edmore/esp/neuron"
	"github.com/edmore/esp/random"

	//	"fmt"
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
	//	fmt.Println(s.Neurons[i].Fitness, s.Neurons[i].Trials, s.Neurons[j].Fitness, s.Neurons[j].Trials)
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

// Mate neurons in population
func (p *Population) Mate() {
	var mate int
	for i := 0; i < p.Numbreed; i++ {
		// Find mate
		if i == 0 {
			mate = rand.Int() % p.Numbreed
		} else {
			mate = rand.Int() % i
		}
		// replace lower half of population
		childIndex1 := p.NumIndividuals - (1 + (i * 2))
		childIndex2 := p.NumIndividuals - (2 + (i * 2))
		onePointCrossover(p.Individuals[i], p.Individuals[mate], p.Individuals[childIndex1], p.Individuals[childIndex2])
	}
}

// Mate and replace the lower ranking half of the population
func onePointCrossover(parent1 *neuron.Neuron, parent2 *neuron.Neuron, child1 *neuron.Neuron, child2 *neuron.Neuron) {
	crosspoint := rand.Int() % len(parent1.Weight) // random crossover point
	for i := 0; i < len(parent1.Weight); i++ {
		child1.Weight[i] = parent2.Weight[i]
		child2.Weight[i] = parent1.Weight[i]
	}
	// update the parent fields
	child1.Parent1 = parent1.Id
	child1.Parent2 = parent2.Id
	child2.Parent1 = parent1.Id
	child2.Parent2 = parent2.Id
	// reset the fitness and trials
	child1.ResetFitness()
	child2.ResetFitness()
	// exchange chromosomal segments
	for j := 0; j < crosspoint; j++ {
		temp := child1.Weight[j]
		child1.Weight[j] = child2.Weight[j]
		child2.Weight[j] = temp
	}
}

// Mutate neurons in population
func (p *Population) Mutate(m float32) {
	for i := p.Numbreed * 2; i < p.NumIndividuals; i++ {
		if rand.Float32() < m {
			mutationIndex := rand.Int() % p.GeneSize
			p.Individuals[i].Weight[mutationIndex] = p.Individuals[i].Weight[mutationIndex] + random.Cauchy(0.3)
		}
	}
}
