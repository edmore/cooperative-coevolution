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
	"time"
)

type Neurons []*neuron.Neuron

// for sorting neurons by average fitness
func (s Neurons) Len() int      { return len(s) }
func (s Neurons) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByAvgFitness implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Neurons value.
type ByAvgFitness struct{ Neurons }

// sort in descending order - largest first
func (s ByAvgFitness) Less(i, j int) bool {
	//prevent division by zero
	divisor1 := s.Neurons[i].Trials
	divisor2 := s.Neurons[j].Trials
	if divisor1 == 0 {
		divisor1 = 1
	}
	if divisor2 == 0 {
		divisor2 = 1
	}
	//	fmt.Println(s.Neurons[i].Fitness, s.Neurons[i].Trials, s.Neurons[j].Fitness, s.Neurons[j].Trials)
	return (s.Neurons[i].Fitness / divisor1) > (s.Neurons[j].Fitness / divisor2)
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
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	index := r.Int() % p.NumIndividuals
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
			mate = r.Int() % p.Numbreed
		} else {
			mate = r.Int() % i
		}
		// replace lower half of population
		childIndex1 := p.NumIndividuals - (1 + (i * 2))
		childIndex2 := p.NumIndividuals - (2 + (i * 2))
		onePointCrossover(p.Individuals[i], p.Individuals[mate], p.Individuals[childIndex1], p.Individuals[childIndex2])
	}
}

// Mate and replace the lower ranking half of the population
func onePointCrossover(parent1 *neuron.Neuron, parent2 *neuron.Neuron, child1 *neuron.Neuron, child2 *neuron.Neuron) {
	crosspoint := r.Int() % len(parent1.Weight) // random crossover point
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
		if r.Float32() < m {
			mutationIndex := r.Int() % p.GeneSize
			p.Individuals[i].Weight[mutationIndex] = p.Individuals[i].Weight[mutationIndex] + random.Cauchy(0.3)
		}
	}
}

// Grow the neuron weights for each individual in the population
func (p *Population) GrowIndividuals() {
	for i := 0; i < p.NumIndividuals; i++ {
		tmp := []float64{1.0}
		p.Individuals[i].Weight = append(p.Individuals[i].Weight, tmp...)
	}
}
