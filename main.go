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
	subpops     []*population.Population
	predSubpops [][]*population.Population
	world       environment.Gridworld
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
	pred        = flag.Int("pred", 3, "predators")
)

type TempState struct {
	PredatorX []int // x position(s) of the predator(s)
	PredatorY []int // y position(s) of the predator(s)
	PreyX     int   // x position of the prey
	PreyY     int   // y position of the prey
}

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
	var state environment.State

	var tempState TempState
	states := make([]TempState, 0)

	state = *e.GetState()
	world = *e.GetWorld()

	nearestDistance := 0
	nearestPredator := 0
	currentDistance := 0

	// calculate average INITIAL distance
	for p := 0; p < *pred; p++ {
		average_initial_distance = average_initial_distance + calculateDistance(state.PredatorX[p], state.PredatorY[p], state.PreyX, state.PreyY)
	}
	average_initial_distance = average_initial_distance / *pred

	for !e.Caught() && steps < maxSteps {

		// find the nearest predator
		for p := 0; p < *pred; p++ {
			currentDistance = calculateDistance(state.PredatorX[p], state.PredatorY[p], state.PreyX, state.PreyY)
			if currentDistance < nearestDistance {
				nearestDistance = currentDistance
				nearestPredator = p
			}
		}
		// Proceed to next state ...

		// Perform prey action
		e.PerformPreyAction(nearestPredator)

		// Perform each predator action
		for key, predator := range team {
			input[0] = float64(state.PreyX)
			input[1] = float64(state.PreyY)

			out := predator.Activate(input, output)
			e.PerformPredatorAction(key, out)
		}
		state = *e.GetState()
		steps++

		// push tempState into the states slice : need to avoid referencing the slice address and only the last state being present
		tempState = *new(TempState)
		var tempPredatorY []int
		var tempPredatorX []int
		for i := 0; i < len(state.PredatorX); i++ {
			tempPredatorX = append(tempPredatorX, state.PredatorX[i])
			tempPredatorY = append(tempPredatorY, state.PredatorY[i])
		}

		tempState = TempState{PredatorX: tempPredatorX, PredatorY: tempPredatorY, PreyX: state.PreyX, PreyY: state.PreyY}
		states = append(states, tempState)
	}

	if e.Caught() {
		if *simulation == true {
			//fmt.Println("Steps ", steps)
			// TODO - You need a clause here to say write to file if prey is caught; so you have a simulation that demonstrates a capture.
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

	// calculate fitness - which is average FINAL distance from the prey
	for p := 0; p < *pred; p++ {
		average_final_distance = average_final_distance + calculateDistance(state.PredatorX[p], state.PredatorY[p], state.PreyX, state.PreyY)
	}
	average_final_distance = average_final_distance / *pred

	if !e.Caught() {
		fitness = (average_initial_distance - average_final_distance) / 10
	} else {
		// best case fitness would be like 20. where both predators capture the prey at the same time
		fitness = (200 - average_final_distance) / 10
		catches++
	}

	// award fitness score to team
	for _, predator := range team {
		predator.SetFitness(fitness)
		predator.SetNeuronFitness()
	}
	return team
}

// Calculate Manhattan Distance
func calculateDistance(predX int, predY int, preyX int, preyY int) int {
	distanceX := 0.0
	distanceY := 0.0

	distanceX = math.Abs(float64(predX - preyX))
	if distanceX > float64(world.Length/2) {
		distanceX = float64(world.Length) - distanceX
	}

	distanceY = math.Abs(float64(predY - preyY))
	if distanceY > float64(world.Height/2) {
		distanceY = float64(world.Height) - distanceY
	}

	return int(distanceX + distanceY)
}

var catches int

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

	var mutationRate float32 = 0.4

	fmt.Printf("Number of inputs (i) is %v.\n", *i)
	fmt.Printf("Number of hidden units (h) is %v.\n", *h)
	fmt.Printf("Number of output(s) is %v.\n", *o)
	fmt.Printf("Number of individuals per population (n) is %v.\n", *n)
	fmt.Printf("Max generations is %v.\n", *maxGens)
	fmt.Printf("Mutation Rate is set at %v.\n", mutationRate)
	fmt.Printf("Burst mutate after %v constant generations (b).\n", *b)

	fmt.Printf("Number of predators is %v.\n", *pred)

	performanceQueue := make([]int, *b)
	bestFitness := 0
	generations := 0

	fmt.Println("Number of Logical CPUs on machine ", runtime.NumCPU())
	defaultCPU := runtime.GOMAXPROCS(0)
	fmt.Println("DefaultCPU(s) ", defaultCPU)
	hiddenUnits := *h

	// INITIALIZATION
	// TODO - work out whether using the network genesize is the best way to do this
	// You have to remember that if you have 3 predators and number of hidden units is 3 that means 9 subpops
	for p := 0; p < *pred; p++ {
		subpops = initialize(hiddenUnits, *n, network.NewFeedForward(*i, hiddenUnits, *o, false).GeneSize)
		// predator subpopulations - a multidimensional array of subpopulations
		predSubpops = append(predSubpops, subpops)
	}

	numTrials := 10 * *n
	// probably need to terminate when the prey has been caught at least 50% (or whatever) of the evaluations by a particular team
	// or based on the average distance (fitness) : selection of the optimal distance from the prey; but this might be harder
	for generations < *maxGens && catches != numTrials {
		// Reset catches
		catches = 0
		// EVALUATION
		// TODO - fix so it is numTrials
		for x := 0; x < numTrials; x++ {
			// Build the team of predators
			// [[p,p,p], [p,p,p]....] - predator subpops
			// has to be a unique network
			var team []network.Network
			for f := 0; f < *pred; f++ {
				feedForward := network.NewFeedForward(*i, hiddenUnits, *o, false)
				feedForward.Create(predSubpops[f])
				team = append(team, feedForward)
			}
			// Evaluate the team in the environment(e)
			e := environment.NewPredatorPrey()
			e.Reset(*pred)

			// TODO : Fix the logic below
			// might need to make the prey starting points random : original code has 9 evaluations and finds average
			// then again remember the networks are different i.e. different random weights
			t := evaluate(e, team)

			if t[0].GetFitness() > bestFitness {
				bestFitness = t[0].GetFitness()
				bestTeam = t

				for i := 0; i < len(bestTeam); i++ {
					bestTeam[i].Tag()
				}
			}
		}

		fmt.Printf("Generation %v, best fitness is %v, catches is %v\n", generations, bestFitness, catches)
		performanceQueue = append(performanceQueue, bestFitness)

		// RECOMBINATION - sort neurons, mate and mutate
		for _, predatorPops := range predSubpops {
			for _, subpop := range predatorPops {
				// Sort neurons in each subpopulation
				subpop.SortNeurons()
				// Mate top quartile of neurons in each population
				subpop.Mate()
				// Mutate lower half of population
				subpop.Mutate(mutationRate)
			}
		}
		generations++

	}
}
