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
	"math"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	bestTeam    []network.Network
	ch          = make(chan network.Network)
	chans       = make([]chan network.Network, 0)
	subpops     []*population.Population
	predSubpops []*population.Population
	world       *environment.Gridworld
)

// Flags
var (
	simulation  = flag.Bool("sim", false, "simulate best network on task")
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to file")
	h           = flag.Int("h", 10, "number of hidden units / subpopulations")
	n           = flag.Int("n", 100, "number of individuals per subpopulation")
	i           = flag.Int("i", 2, " number of inputs")
	o           = flag.Int("o", 5, "number of outputs")
	b           = flag.Int("b", 10, "number of generations before burst mutation")
	maxGens     = flag.Int("maxGens", 100000, "maximum generations")
	goalFitness = flag.Int("goalFitness", 100000, "goal fitness")
	p           = flag.Int("pred", 3, "predators")
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

// Evaluate the team of networks (predators) in the trial environment
func evaluate(e environment.Environment, team []network.Network) []network.Network {
	fitness := 0
	steps := 0
	maxSteps := 150
	average_initial_distance := 0
	average_final_distance := 0

	input := make([]float64, team[0].GetTotalInputs())   // position of the prey
	output := make([]float64, team[0].GetTotalOutputs()) // N,S,E,W,Stay
	var state *environment.State
	states := make([]environment.State, 0)

	state = e.GetState()
	world = e.GetWorld()

	// calculate average INITIAL distance
	for p := 0; p < numPreds; p++ {
		average_initial_distance = average_initial_distance + calculateDistance(e, state.PredatorX[p], state.PredatorY[p], state.PreyX, state.PreyY)
	}
	average_initial_distance = average_initial_distance / numPreds

	for !e.Caught() && steps < maxSteps {
		// push state into states slice
		states = append(states, *state)
		// Proceed to next state ...

		// Perform prey action
		e.PerformPreyAction()

		// Perform each predator action
		for _, predator := range team {
			input[0] = state.PreyX
			input[1] = state.PreyY

			out := predator.Activate(input, output)
			e.PerformPredatorAction(predator, out)
		}
		steps++
		state = e.GetState()
	}

	if *simulation == true {
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
	// calculate fitness - which is average FINAL distance from the prey
	for p := 0; p < numPreds; p++ {
		average_final_distance = average_final_distance + calculateDistance(e, state.PredatorX[p], state.PredatorY[p], state.PreyX, state.PreyY)
	}
	average_final_distance = average_final_distance / numPreds

	if !e.Caught() {
		fitness = (average_initial_distance - average_final_distance) / 10
	} else {
		fitness = (200 - average_final_distance) / 10
	}

	// award fitness score to team
	for _, predator := range team {
		predator.SetFitness(fitness)
		predator.SetNeuronFitness()
	}
	return team
}

// Calculate Manhattan Distance
func calculateDistance(e environment.Environment, predX int, predY int, preyX int, preyY int) {
	distanceX := 0
	distanceY := 0

	distanceX = math.Abs(predX - preyX)
	if distanceX > world.Length/2 {
		distanceX = world.Length - distanceX
	}

	distanceY = math.Abs(predY - preyY)
	if distanceY > world.Height/2 {
		distanceY = world.Height - distanceY
	}

	return (distanceX + distanceY)
}

var numPreds int = *p

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

	fmt.Printf("Number of predators is %v.\n", *p)

	performanceQueue := make([]int, *b)
	bestFitness := 0
	generations := 0
	stagnated = false
	count := 0

	fmt.Println("Number of Logical CPUs on machine ", runtime.NumCPU())
	defaultCPU := runtime.GOMAXPROCS(0)
	fmt.Println("DefaultCPU(s) ", defaultCPU)
	hiddenUnits := *h

	// INITIALIZATION
	// TODO - work out whether using the network genesize is the best way to do this
	// You have to remember that if you have 3 predators and number of hidden units is 3 that means 9 subpops
	for p := 0; p < numPreds; p++ {
		subpops = initialize(hiddenUnits, *n, network.NewFeedForward(*i, hiddenUnits, *o, false).GeneSize)
		// predator subpopulations - a multidimensional array
		predSubpops.append(subpops)
	}

	numTrials := 10 * *n
	for generations < *maxGens {
		// EVALUATION
		for x := 0; x < numTrials; x++ {
			// Build the team of predators
			// [[p,p,p], [p,p,p]....]
			// has to be a unique network
			var team []network.Network
			for i := 0; i < len(predSubpops); i++ {
				feedForward := network.NewFeedForward(*i, hiddenUnits, *o, false)
				feedForward.Create(predSubpops[i])
				team.append(feedForward)
			}
			// Evaluate the team in the environment(e)
			e := environment.NewPredatorPrey()
			e.Reset()
			t := evaluate(e, team)
			if t[0].GetFitness() > bestFitness {
				bestFitness = t[0].GetFitness()
				bestTeam = t
				for i := 0; i < len(bestTeam); i++ {
					bestTeam[i].Tag()
				}
			}
		}

		fmt.Printf("Generation %v, best fitness is %v\n", generations, bestFitness)
		performanceQueue = append(performanceQueue, bestFitness)

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
