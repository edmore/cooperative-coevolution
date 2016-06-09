// +build ignore

package main

import (
	"github.com/edmore/cooperative-coevolution/environment"
	"github.com/edmore/cooperative-coevolution/network"
	"github.com/edmore/cooperative-coevolution/neuron"
)

var ch = make(chan []*neuron.Neuron)

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network) {
	// loop while within bounds
	// e.PerformAction(n.Activate())
	// fitness++

	// award fitness score to network
	// define setter method for network fitness

	// add the fitness score to cumulative fitness of
	// neurons that participated in trial.
	// beware of race conditions when adding the fitness
	// to each neuron needs to be synchronized.
	ch <- n.GetHiddenUnits()
}
