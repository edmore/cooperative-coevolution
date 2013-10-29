package main

import (
	"fmt"
	"github.com/edmore/esp/network"
	"github.com/edmore/esp/population"
	"reflect"
	"runtime"
	"time"
)

func initialize(h int, n int, s int) []*population.Population {
	var pops []*population.Population // population pool
	ch := make(chan *population.Population)

	for i := 0; i < h; i++ {
		go func() {
			fmt.Println("Creating subpopulation ...")
			p := population.NewPopulation(n, s)
			p.Create()
			ch <- p
		}()
	}

	for {
		select {
		case pop := <-ch:
			pops = append(pops, pop)
			fmt.Println("Subpopulation initialized.")
			if len(pops) == h {
				return pops
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return pops
}

func evaluate(n network.Network) {
	fmt.Println(reflect.TypeOf(n))
}

func main() {
	nCPU := runtime.NumCPU()
	cCPU := 2
	runtime.GOMAXPROCS(cCPU)
	fmt.Println("Number of CPUs available: ", nCPU)
	fmt.Println("Number of CPUs currently in use: ", cCPU)

	var (
		h int // number of hidden units / subpopulations
		n int // number of neuron chromosomes per subpopulation
		i int // number of inputs
		o int // number of outputs
	)

	fmt.Printf("Please enter the number of hidden units : ")
	fmt.Scanf("%d", &h)
	fmt.Printf("Please enter the number of neuron chromosomes per population : ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Please enter the number of inputs : ")
	fmt.Scanf("%d", &i)
	fmt.Printf("Please enter the number of outputs : ")
	fmt.Scanf("%d", &o)

	feedForward := network.NewFeedForward(i, h, o, 1)

	// Initialization
	subpops := initialize(h, n, feedForward.GeneSize)
	fmt.Println(subpops[0])
	fmt.Println(subpops[0].Neurons[0])

	// Evaluation
	feedForward.Create(subpops)
	fmt.Println("First Hidden Unit Id:", feedForward.HiddenUnits[0].Id)
	fmt.Println("Second Hidden Unit Id:", feedForward.HiddenUnits[1].Id)
	evaluate(feedForward)
}
