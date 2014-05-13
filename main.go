package main

import (
	"fmt"
	"time"

	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
	"github.com/edmore/esp/population"
)

// Evaluator interface
type Evaluator interface {
	evaluate(environment.Environment, network.Network)
}

var (
	maxFitness  int = 100000 // the maximum fitness in time steps
	bestNetwork network.Network
)

// Initialize subpopulations
func initialize(h int, n int, s int) []*population.Population {
	var pops []*population.Population // population pool

	for i := 0; i < h; i++ {
		p := population.NewPopulation(n, s)
		p.Create()
		pops = append(pops, p)
	}
	return pops
}

func main() {
	var (
		h              int // number of hidden units / subpopulations
		n              int // number of neuron chromosomes per subpopulation
		i              int // number of inputs
		o              int // number of outputs
		maxGenerations int // maximum generations
	)

	//	fmt.Printf("Please enter the number of inputs : ")
	//      fmt.Scanf("%d", &i)
	fmt.Println("Number of inputs is 6 (Markov)")
	fmt.Printf("Please enter the number of hidden units : ")
	fmt.Scanf("%d", &h)
	fmt.Printf("Please enter the number of outputs : ")
	fmt.Scanf("%d", &o)
	fmt.Printf("Please enter the number of neuron chromosomes per population : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Please enter the max generations : ")
	fmt.Scanf("%d", &maxGenerations)

	bestFitness := 0
	generations := 0
	i = 6 // Double Pole balancing Task (Markov)

	// INITIALIZATION
	// TODO - work out whether using the network genesize is the best way to do this
	subpops := initialize(h, n, network.NewFeedForward(i, h, o, true).GeneSize)

	for bestFitness < maxFitness && generations < maxGenerations {
		generations++
		numTrials := 10 * n
		// EVALUATION
		for x := 0; x < numTrials; x++ {
			// Build the network
			feedForward := network.NewFeedForward(i, h, o, true)
			feedForward.Create(subpops)
			// Evaluate the network in the environment(e)
			e := environment.NewCartpole()
			e.Reset()
			go evaluate(e, feedForward)
		}
	ForSelect:
		for {
			select {
			case network := <-ch:
				network.SetNeuronFitness()
				if network.GetFitness() > bestFitness {
					bestFitness = network.GetFitness()
					bestNetwork = network
				}
			case <-time.After(500 * time.Millisecond):
				break ForSelect
			}
		}

		// CHECK STAGNATION
		// if bestFitness has not improved in b generations
		// if this is the second consecutive time
		// then ADAPT-NETWORK-SIZE()
		// else BURST_MUTATE()

		// RECOMBINATION - sort neurons, mate and mutate
		// TODO - optional to recombine if you just burst mutated.
		for _, subpop := range subpops {
			// Sort neurons in each subpopulation
			subpop.SortNeurons()
			// Mate top quartile of neurons in each population
			// The top quartile neurons need to be reset as well - fitness and trials
			subpop.Mate()
			// Mutate
		}
	}
}
