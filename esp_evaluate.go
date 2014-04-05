package main

import (
	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
)

var ch = make(chan network.Network)

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network) {
	// loop while within bounds
	// e.PerformAction(n.Activate())
	// fitness++
	fitness := 0
	for e.WithinTrackBounds() && e.WithinAngleBounds() {
		// inputs has to be a slice for a general framework
		inputs := e.GetState()

		output := n.Activate(inputs)
		e.PerformAction(output)
		fitness++
	}
	// award fitness score to network
	// define setter method for network fitness

	// add the fitness score to cumulative fitness of
	// neurons that participated in trial.
	// beware of race conditions when adding the fitness
	// to each neuron needs to be synchronized.
	ch <- n
}
