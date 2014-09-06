package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

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
)

// Flags
var (
	simulation  = flag.Bool("sim", false, "simulate best network on task")
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
		input[0] = state.X / 4.8
		input[1] = state.XDot / 2
		input[2] = state.Theta1 / 0.52
		input[3] = state.Theta2 / 0.52
		input[4] = state.ThetaDot1 / 2
		input[5] = state.ThetaDot2 / 2
		if n.HasBias() {
			input[6] = 0.5 // bias
		}
		out := n.Activate(input, output)
		e.PerformAction(out[0])
		lesionedFitness++
	}
	return lesionedFitness
}

// Run a split of evaluations
func splitEvals(numTrials int, numCPU int, i int, h int, o int, subpops []*population.Population, c chan network.Network) {
	var phaseBestNetwork network.Network
	phaseBestFitness := 0

	for x := 0; x < (numTrials / numCPU); x++ {
		// Build the network
		feedForward := network.NewFeedForward(i, h, o, true)
		feedForward.Create(subpops)
		// Evaluate the network in the environment(e)
		e := environment.NewCartpole()
		e.Reset()
		go evaluate(e, feedForward, c)
	}
	for x := 0; x < (numTrials / numCPU); x++ {
		network := <-c
		if network.GetFitness() > phaseBestFitness {
			phaseBestFitness = network.GetFitness()
			phaseBestNetwork = network
		}
	}
	ch <- phaseBestNetwork
}

// Create the json file for use in simulation
func createJSON() {
	simulatedFitness := 0
	input := make([]float64, bestNetwork.GetTotalInputs())
	output := make([]float64, bestNetwork.GetTotalOutputs())
	var state *environment.State
	states := make([]environment.State, 0)

	simEnviron := environment.NewCartpole()
	simEnviron.Reset()

	for simulatedFitness < *goalFitness {
		state = simEnviron.GetState()
		// push state into states slice
		states = append(states, *state)
		// Proceed to next state
		input[0] = state.X / 4.8
		input[1] = state.XDot / 2
		input[2] = state.Theta1 / 0.52
		input[3] = state.Theta2 / 0.52
		input[4] = state.ThetaDot1 / 2
		input[5] = state.ThetaDot2 / 2
		if bestNetwork.HasBias() {
			input[6] = 0.5 // bias
		}
		out := bestNetwork.Activate(input, output)
		simEnviron.PerformAction(out[0])
		simulatedFitness++
	}
	// write the states to a json file
	b, err := json.Marshal(states)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ioutil.WriteFile("simulation/processingjs/json/states.json", b, 0644)
	if err != nil {
		panic(err)
	}
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
	subpops := initialize(hiddenUnits, *n, network.NewFeedForward(*i, hiddenUnits, *o, true).GeneSize)

	numTrials := 10 * *n
	for bestFitness < *goalFitness && generations < *maxGens {
		// EVALUATION
		runtime.GOMAXPROCS(numCPU)
		// Distribute a split of evaluations over multiple cores/CPUs
		for y := 0; y < numCPU; y++ {
			chans = append(chans, make(chan network.Network))
			go splitEvals(numTrials, numCPU, *i, hiddenUnits, *o, subpops, chans[y])
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
						p := population.NewPopulation(*n, network.NewFeedForward(*i, hiddenUnits, *o, true).GeneSize)
						p.Create()
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

	// Start Webserver for simulation
	if *simulation != false {
		// write states to file
		createJSON()
		log.Println("Starting simulation server ...")
		log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("simulation/processingjs/"))))
	}
}
