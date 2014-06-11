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

func evaluateLesioned(e environment.Environment, n network.Network) int {
	lesionedFitness := 0
	for e.WithinTrackBounds() && e.WithinAngleBounds() {
		state := e.GetState()
		input := make([]float64, n.GetTotalInputs())
		input[0] = state.X / 4.8
		input[1] = state.XDot / 2
		input[2] = state.Theta1 / 0.52
		input[3] = state.Theta2 / 0.52
		input[4] = state.ThetaDot1 / 2
		input[5] = state.ThetaDot2 / 2
		if n.HasBias() {
			input[6] = 0.5 // bias
		}
		output := n.Activate(input)
		e.PerformAction(output[0])
		lesionedFitness++
	}
	return lesionedFitness
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
					bestNetwork.Tag()
				}
			case <-time.After(500 * time.Millisecond):
				break ForSelect
			}
		}
		fmt.Printf("Generation %v best fitness is %v\n", generations, bestFitness)
		performanceQueue = append(performanceQueue, bestFitness)

		// CHECK STAGNATION
		// if bestFitness has not improved in b generations
		//   if fitness has not improved after two(2) burst mutations
		//   then ADAPT-NETWORK-SIZE()
		//   else BURST_MUTATE()
		if performanceQueue[b+generations] == performanceQueue[generations] {
			stagnated = true
			if count == 2 {
				fmt.Println("Adapting network size ...")
				for item, neuron := range bestNetwork.GetHiddenUnits() {
					neuron.Lesioned = true
					lesionedEnviron := environment.NewCartpole()
					lesionedEnviron.Reset()

					lesionedFitness := evaluateLesioned(lesionedEnviron, bestNetwork)
					fmt.Println("Lesioned Fitness: ", lesionedFitness)
					// TODO : Determine the threshold to use
					threshold := 1

					if lesionedFitness > (bestFitness * threshold) {
						// delete subpopulation to subpops
						//decrement h
						subpops = append(subpops[:item], subpops[item+1:]...)
						h--
						fmt.Println("Subpopulations decreased to ", h)
					} else {
						neuron.Lesioned = false
					}
				}
				// if no neuron is removed
				// add a new population to subpops
				// increment h
				if len(subpops) == h {
					h++
					fmt.Println("Subpopulations increased to ", h)
					p := population.NewPopulation(n, network.NewFeedForward(i, h, o, true).GeneSize)
					p.Create()
					subpops = append(subpops, p)
				}
				count = 0
			} else {
				// Do not burst mutate just after you adapt the network size
				if len(bestNetwork.GetHiddenUnits()) == h {
					fmt.Println("Burst Mutate ...")

					for index, subpop := range subpops {
						for _, neuron := range subpop.Individuals {
							neuron.Perturb(bestNetwork.GetHiddenUnits()[index])
						}
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
