package main

import (
	"fmt"
	"github.com/edmore/esp/population"
	"runtime"
	"time"
)

func initialize(h int, n int) []*population.Population {
	var pops []*population.Population // population pool
	ch := make(chan *population.Population)

	for i := 0; i < h; i++ {
		go func() {
			fmt.Println("Creating subpopulation ...")
			p := population.NewPopulation(n)
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

func main() {
	nCPU := runtime.NumCPU()
	cCPU := 2
	runtime.GOMAXPROCS(cCPU)
	fmt.Println("Number of CPUs available: ", nCPU)
	fmt.Println("Number of CPUs currently in use: ", cCPU)

	var h int // number of hidden units / subpopulations
	var n int // number of neuron chromosomes per subpopulation

	fmt.Printf("Please enter the number of hidden units : ")
	fmt.Scanf("%d", &h)
	fmt.Printf("Please enter the number of neuron chromosomes per population : ")
	fmt.Scanf("%d", &n)

	subpops := initialize(h, n)
	fmt.Println(subpops)
}
