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

var maxFitness float64 = 100000.0 // the maximum fitness in time steps

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
		h int // number of hidden units / subpopulations
		n int // number of neuron chromosomes per subpopulation
		i int // number of inputs
		o int // number of outputs
	)

	fmt.Printf("Please enter the number of inputs : ")
	fmt.Scanf("%d", &i)
	fmt.Printf("Please enter the number of hidden units : ")
	fmt.Scanf("%d", &h)
	fmt.Printf("Please enter the number of outputs : ")
	fmt.Scanf("%d", &o)
	fmt.Printf("Please enter the number of neuron chromosomes per population : ")
	fmt.Scanf("%d", &n)

	bestFitness := 0.0

	for bestFitness < maxFitness {
		// INITIALIZATION
		// TODO - work out whether using the network genesize is the best way to do this
		subpops := initialize(h, n, network.NewFeedForward(i, h, o, 1).GeneSize)
		//fmt.Println(subpops)

		numTrials := 10 * n
		// EVALUATION
		for x := 0; x < numTrials; x++ {
			// Build the network
			feedForward := network.NewFeedForward(i, h, o, 1)
			feedForward.Create(subpops)

			// Evaluate the network in the environment(e)
			e := environment.NewCartpole()
			e.Reset()
			go evaluate(e, feedForward)
		}
		// block to add fitness to each neuron and ...
		// also save bestFitness (and probably best-network)
		// if fitness > bestFitness ... save bestFitness
		for {
			select {
			case n := <-ch:
				// You can define a setter method for setting the neuron fitness
				fmt.Println(n.GetHiddenUnits())
			case <-time.After(50 * time.Millisecond):
				return
			}
		}
		bestFitness = maxFitness

		// CHECK STAGNATION
		// if bestFitness has not improved in b generations
		// if this is the second consecutive time
		// then ADAPT-NETWORK-SIZE()
		// else BURST_MUTATE()

		// RECOMBINATION
		// sort neurons - mate and mutate
	}
}
