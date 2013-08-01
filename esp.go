package main

import (
	"fmt"
	"github.com/edmore/esp/population"
	"sync"
	"time"
)

func main() {
	var pops [10]*population.Population
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			p := population.NewPopulation(1000)
			p.Create()
			pops[i] = p
			wg.Done()
		}(i)
	}
	wg.Wait()
	// do something else while you wait ...
	time.Sleep(15)
	// after some time pops is available for use
	fmt.Println(pops)
}
