package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
	"github.com/edmore/esp/population"
	"io/ioutil"
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
)

// Flags
var (
	simulation  = flag.Bool("sim", false, "simulate best network on task")
	markov      = flag.Bool("markov", false, "Markov or Non-Markov task")
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to file")
	h           = flag.Int("h", 10, "number of hidden units / subpopulations")
	n           = flag.Int("n", 20, "number of individuals per subpopulation")
	i           = flag.Int("i", 6, " number of inputs")
	o           = flag.Int("o", 1, "number of outputs")
	maxGens     = flag.Int("maxGens", 100000, "maximum generations")
	goalFitness = flag.Int("goalFitness", 100000, "goal fitness")
)

// Initialize subpopulations
func initialize(n int, s int) []*population.Population {
	var pops []*population.Population // population pool

	p := population.NewPopulation(n, s)
	p.Create()
	pops = append(pops, p)

	return pops
}

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network) network.Network {
	fitness := 0
	input := make([]float64, n.GetTotalInputs())
	output := make([]float64, n.GetTotalOutputs())
	var state *environment.State
	states := make([]environment.State, 0)

	for e.WithinTrackBounds() && e.WithinAngleBounds() && fitness < *goalFitness {
		state = e.GetState()
		// push state into states slice
		states = append(states, *state)
		// Proceed to next state
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
		fitness++
	}

	if *simulation == true {
		if fitness == *goalFitness {
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
	}
	// award fitness score to network
	n.SetFitness(fitness)
	n.SetNeuronFitness()
	return n
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

	bestFitness := 0
	generations := 0
	stagnated = false

	fmt.Println("Number of Logical CPUs on machine ", runtime.NumCPU())
	defaultCPU := runtime.GOMAXPROCS(0)
	fmt.Println("DefaultCPU(s) ", defaultCPU)
	hiddenUnits := *h

	// INITIALIZATION
	// TODO - work out whether using the network genesize is the best way to do this
	if *markov == true {
		subpops = initialize(*n, network.NewFeedForward(*i, hiddenUnits, *o, true).GeneSize)
	} else {
		subpops = initialize(*n, network.NewRecurrent(*i, hiddenUnits, *o, true).GeneSize)
	}

	numTrials := 10 * *n
	for bestFitness < *goalFitness && generations < *maxGens {
		// EVALUATION
		for x := 0; x < numTrials; x++ {
			if *markov == true {
				// Build the network
				feedForward := network.NewFeedForward(*i, hiddenUnits, *o, true)
				feedForward.Create(subpops)
				// Evaluate the network in the environment(e)
				e := environment.NewCartpole()
				e.Reset()
				n := evaluate(e, feedForward)
				if n.GetFitness() > bestFitness {
					bestFitness = n.GetFitness()
					bestNetwork = n
					bestNetwork.Tag()
				}

			} else {
				// Build the network
				recurrent := network.NewRecurrent(*i, hiddenUnits, *o, true)
				recurrent.Create(subpops)
				// Evaluate the network in the environment(e)
				e := environment.NewCartpole()
				e.Reset()
				n := evaluate(e, recurrent)
				if n.GetFitness() > bestFitness {
					bestFitness = n.GetFitness()
					bestNetwork = n
					bestNetwork.Tag()
				}

			}
		}

		fmt.Printf("Generation %v, best fitness is %v\n", generations, bestFitness)

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
