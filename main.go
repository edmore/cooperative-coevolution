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
	goalFitness int = 100000 // the goal fitness in time steps
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
		h              int     // number of hidden units / subpopulations
		n              int     // number of individuals per subpopulation
		i              int     // number of inputs
		o              int     // number of outputs
		b              int     // number of generations before burst mutation
		maxGenerations int     // maximum generations
		mutationRate   float32 // rate of mutation
		stagnated      bool
	)

	fmt.Println("Number of inputs is 6 (Markov)")
	fmt.Printf("Please enter the number of hidden units : ")
	fmt.Scanf("%d", &h)
	fmt.Printf("Please enter the number of output(s) : ")
	fmt.Scanf("%d", &o)
	fmt.Printf("Please enter the number of individuals per population : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Please enter the max generations : ")
	fmt.Scanf("%d", &maxGenerations)
	fmt.Printf("Mutation Rate is set at 0.4.\n")
	fmt.Printf("Burst mutate after how many constant generations? : ")
	fmt.Scanf("%d", &b)

	performanceQueue := make([]int, b)
	bestFitness := 0
	generations := 0
	i = 6 // Double Pole balancing Task (Markov)
	mutationRate = 0.4
	stagnated = false
	count := 0

	// INITIALIZATION
	// TODO - work out whether using the network genesize is the best way to do this
	subpops := initialize(h, n, network.NewFeedForward(i, h, o, true).GeneSize)

	for bestFitness < goalFitness && generations < maxGenerations {

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
		fmt.Printf("Generation %v best fitness is %v\n", generations, bestFitness)
		performanceQueue = append(performanceQueue, bestFitness)
		fmt.Println(performanceQueue)

		// CHECK STAGNATION
		// if bestFitness has not improved in b generations
		//   if fitness has not improved after two(2) burst mutations
		//   then ADAPT-NETWORK-SIZE()
		//   else BURST_MUTATE()
		if performanceQueue[b+generations] == performanceQueue[generations] {
			if count == 2 {
				fmt.Println("Adapting network size ...")
				count = 0
			} else {
				stagnated = true
				fmt.Println("Burst Mutate ...")
				for index, subpop := range subpops {
					for _, neuron := range subpop.Individuals {
						neuron.Perturb(bestNetwork.GetHiddenUnits()[index])
					}
				}
				count++
			}
		}

		// RECOMBINATION - sort neurons, mate and mutate
		if stagnated == false {
			for _, subpop := range subpops {
				// Sort neurons in each subpopulation
				subpop.SortNeurons()
				// Mate top quartile of neurons in each population
				subpop.Mate()
				// Mutate lower half of population
				subpop.Mutate(mutationRate)
			}
		}
		// reset stagnation
		stagnated = false
		// Add fitness to performance queue
		generations++
	}
}
