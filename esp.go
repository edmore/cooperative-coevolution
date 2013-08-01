package main

import (
	"fmt"
	"github.com/edmore/esp/population"
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
			fmt.Println("Subpopulation initialized.")
			pops = append(pops, pop)
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
	var h int // number of hidden units / subpopulations
	var n int // number of neuron chromosomes per subpopulation

	fmt.Printf("Please enter the number of hidden units : ")
	fmt.Scanf("%d", &h)
	fmt.Printf("Please enter the number of neuron chromosomes per population : ")
	fmt.Scanf("%d", &n)

	subpops := initialize(h, n)
	fmt.Println(subpops)
}
