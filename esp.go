package main

import (
	"fmt"

	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
	"github.com/edmore/esp/population"
	"runtime"

//	"time"
)

var maxFitness float64 = 100000.0 // the maximum fitness in time steps

func initialize(h int, n int, s int) []*population.Population {
	nCPU := runtime.NumCPU()
	cCPU := nCPU - 3
	// add more threads to the goroutine schedulers thread pool (default is 1)
	runtime.GOMAXPROCS(cCPU)
	//fmt.Println("Number of Logical CPUs currently in use: ", cCPU)
	//fmt.Println("Number of Logical CPUs available: ", nCPU)
	var pops []*population.Population // population pool

	// Using goroutines and channels - overhead due to use of channels
	//ch := make(chan *population.Population)
	/*
		for i := 0; i < h; i++ {
			go func() {
				//fmt.Println("Creating subpopulation ...")
				p := population.NewPopulation(n, s)
				p.Create()
				ch <- p
			}()
		}

		for {
			select {
			case pop := <-ch:
				pops = append(pops, pop)
				//fmt.Println("Subpopulation initialized.")
				if len(pops) == h {
					return pops
				}
			case <-time.After(50 * time.Millisecond):
				fmt.Printf(".")
			}
		}
	*/

	// Using a basic "for" loop - with "blocking" - actually faster
	for i := 0; i < h; i++ {
		//fmt.Println("Creating subpopulation ...")
		p := population.NewPopulation(n, s)
		p.Create()
		pops = append(pops, p)
	}
	return pops
}

func evaluate(e environment.Environment, n network.Network) float64 {
	// award fitness score to network
	// add the fitness score to cumulative fitness of neurons that participated in trial
	return maxFitness
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

	numTrials := 10 * n

	for x := 0; x < numTrials; x++ {
		feedForward := network.NewFeedForward(i, h, o, 1)

		// Initialization
		subpops := initialize(h, n, feedForward.GeneSize)

		// Building the network
		feedForward.Create(subpops)

		// Evaluation of the network in environment(e)
		e := environment.NewCartpole()
		e.Reset()
		// TODO : Call a goroutine for evaluation
		fmt.Println(evaluate(e, feedForward))
	}
}
