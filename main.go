package main

import (
	"flag"
	"fmt"

	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
	"github.com/edmore/esp/population"

	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	bestNetwork network.Network
	ch          = make(chan network.Network)
	chans       = make([]chan network.Network, 0)
	subpops     []*population.Population
	nets        []network.Network
)

// Flags
var (
	simulation  = flag.Bool("sim", false, "simulate best network on task")
	markov      = flag.Bool("markov", true, "Markov or Non-Markov task")
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to file")
	cpus        = flag.Int("cpus", 1, "number of cpus to use")
	h           = flag.Int("h", 10, "number of hidden units / subpopulations")
	n           = flag.Int("n", 20, "number of individuals per subpopulation")
	i           = flag.Int("i", 6, " number of inputs")
	o           = flag.Int("o", 1, "number of outputs")
	b           = flag.Int("b", 10, "number of generations before burst mutation")
	maxGens     = flag.Int("maxGens", 100000, "maximum generations")
	goalFitness = flag.Int("goalFitness", 100000, "goal fitness")
)

// Initialize subpopulations
func initialize(h int, n int, s int) []*population.Population {
	var pops []*population.Population // population pool

	for w := 0; w < h; w++ {
		p := population.NewPopulation(n, s)
		p.Create()
		pops = append(pops, p)
	}
	return pops
}

// Evaluate a lesioned network
func evaluateLesioned(e environment.Environment, n network.Network) int {
	lesionedFitness := 0
	input := make([]float64, n.GetTotalInputs())
	output := make([]float64, n.GetTotalOutputs())

	for e.WithinTrackBounds() && e.WithinAngleBounds() && lesionedFitness < *goalFitness {
		state := e.GetState()
		if *markov == true {
			input[0] = state.X / 4.8
			input[1] = state.XDot / 2
			input[2] = state.Theta1 / 0.52
			input[3] = state.Theta2 / 0.52
			input[4] = state.ThetaDot1 / 2
			input[5] = state.ThetaDot2 / 2
			if n.HasBias() {
				input[6] = 0.5 // bias
			}
		} else {
			input[0] = state.X / 4.8
			input[1] = state.Theta1 / 0.52
			input[2] = state.Theta2 / 0.52
			if n.HasBias() {
				input[3] = 0.5 // bias
			}
		}
		out := n.Activate(input, output)
		e.PerformAction(out[0])
		lesionedFitness++
	}
	return lesionedFitness
}

// Run a split of evaluations
func splitEvals(split int, nets []network.Network, c chan network.Network) {
	var phaseBestNetwork network.Network
	phaseBestFitness := 0

	for x := 0; x < split; x++ {
		// Evaluate the network in the environment(e)
		e := environment.NewCartpole()
		e.Reset()
		go evaluate(e, nets[x], c)
		//fmt.Println(nets[x])
	}
	for x := 0; x < split; x++ {
		network := <-c
		if network.GetID() == 2 {
			//fmt.Println(network)
		}
		if network.GetFitness() > phaseBestFitness {
			phaseBestFitness = network.GetFitness()
			phaseBestNetwork = network
		}
	}
	fmt.Printf("Core best is %v\n", phaseBestNetwork.GetFitness())
	ch <- phaseBestNetwork
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var (
		stagnated    bool
		mutationRate float32 = 0.4
	)

	// number of inputs for Non-markov Task
	if *markov == false {
		*i = 3
	}

	fmt.Printf("Number of inputs (i) is %v.\n", *i)
	fmt.Printf("Number of hidden units (h) is %v.\n", *h)
	fmt.Printf("Number of output(s) is %v.\n", *o)
	fmt.Printf("Number of individuals per population (n) is %v.\n", *n)
	fmt.Printf("Max generations is %v.\n", *maxGens)
	fmt.Printf("Mutation Rate is set at %v.\n", mutationRate)
	fmt.Printf("Burst mutate after %v constant generations (b).\n", *b)

	performanceQueue := make([]int, *b)
	bestFitness := 0
	generations := 0
	stagnated = false
	count := 0

	fmt.Println("Number of Logical CPUs on machine ", runtime.NumCPU())
	defaultCPU := runtime.GOMAXPROCS(0)
	fmt.Println("DefaultCPU(s) ", defaultCPU)
	numCPU := *cpus
	hiddenUnits := *h
	fmt.Println("CPU(s) in use ", numCPU)
	// INITIALIZATION
	// TODO - work out whether using the network genesize is the best way to do this
	if *markov == true {
		subpops = initialize(hiddenUnits, *n, network.NewFeedForward(*i, hiddenUnits, *o, true).GeneSize)
	} else {
		subpops = initialize(hiddenUnits, *n, network.NewRecurrent(*i, hiddenUnits, *o, true).GeneSize)
	}

	numTrials := 10 * *n
	for bestFitness < *goalFitness && generations < *maxGens {
		// EVALUATION
		// Create networks
		for z := 0; z < numTrials; z++ {
			if *markov == true {
				// Build the network
				feedForward := network.NewFeedForward(*i, hiddenUnits, *o, true)
				feedForward.Create(subpops)
				nets = append(nets, feedForward)
			} else {
				// Build the network
				recurrent := network.NewRecurrent(*i, hiddenUnits, *o, true)
				recurrent.Create(subpops)
				nets = append(nets, recurrent)
			}
		}

		runtime.GOMAXPROCS(numCPU)
		// Distribute a split of evaluations over multiple cores/CPUs
		split := numTrials / numCPU
		start := 0
		end := split
		for y := 0; y < numCPU; y++ {
			fmt.Printf("start %v, end %v\n", start, end)
			chans = append(chans, make(chan network.Network))
			go splitEvals(split, nets[start:end], chans[y])
			start = end
			end = end + split
		}
		for z := 0; z < numCPU; z++ {
			network := <-ch
			network.SetNeuronFitness()
			if network.GetFitness() > bestFitness {
				bestFitness = network.GetFitness()
				bestNetwork = network
				bestNetwork.Tag()
			}
		}
		runtime.GOMAXPROCS(defaultCPU)
		fmt.Printf("Generation %v, best fitness is %v\n", generations, bestFitness)
		performanceQueue = append(performanceQueue, bestFitness)

		// CHECK STAGNATION
		// if bestFitness has not improved in b generations
		//   if fitness has not improved after two(2) burst mutations
		//   then ADAPT-NETWORK-SIZE()
		//   else BURST_MUTATE()
		if len(bestNetwork.GetHiddenUnits()) == hiddenUnits {
			if performanceQueue[*b+generations] == performanceQueue[generations] {
				if count == 2 {
					fmt.Println("Adapting network size ...")
					for item, neuron := range bestNetwork.GetHiddenUnits() {
						neuron.Lesioned = true
						lesionedEnviron := environment.NewCartpole()
						lesionedEnviron.Reset()

						lesionedFitness := evaluateLesioned(lesionedEnviron, bestNetwork)
						fmt.Println("Lesioned Fitness: ", lesionedFitness)

						threshold := 1
						if lesionedFitness > (bestFitness*threshold) && len(bestNetwork.GetHiddenUnits()) == hiddenUnits {
							// delete subpopulation to subpops
							// decrement h
							subpops = append(subpops[:item], subpops[item+1:]...)
							hiddenUnits--
							fmt.Println("Subpopulations decreased to ", hiddenUnits)
						} else {
							neuron.Lesioned = false
						}
					}
					// if no neuron was removed
					// increment h
					// add a new population to subpops
					if len(bestNetwork.GetHiddenUnits()) == hiddenUnits {
						hiddenUnits++
						fmt.Println("Subpopulations increased to ", hiddenUnits)

						var p *population.Population
						if *markov == true {
							p = population.NewPopulation(*n, network.NewFeedForward(*i, hiddenUnits, *o, true).GeneSize)
						} else {
							p = population.NewPopulation(*n, network.NewRecurrent(*i, hiddenUnits, *o, true).GeneSize)
						}
						p.Create()
						// Grow the neuron connection weights in the already existent populations
						if *markov == false {
							for _, subpop := range subpops {
								subpop.GrowIndividuals()
							}
						}
						subpops = append(subpops, p)
					}
					count = 0
				} else {
					fmt.Println("Burst Mutate ...")
					stagnated = true
					for index, subpop := range subpops {
						for _, neuron := range subpop.Individuals {
							neuron.Perturb(bestNetwork.GetHiddenUnits()[index])
						}
					}
					count++
				}
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
		// reset channels
		chans = make([]chan network.Network, 0)
		generations++
	}

}
