package main

import (
	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
)

var ch = make(chan network.Network)

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network) {
	fitness := 0
	for e.WithinTrackBounds() && e.WithinAngleBounds() {
		state := e.GetState()
		input := make([]float64, n.GetTotalInputs())
		input[0] = state.X
		input[1] = state.XDot
		input[2] = state.Theta1
		input[3] = state.Theta2
		input[4] = state.ThetaDot1
		input[5] = state.ThetaDot2

		if n.HasBias() {
			input[6] = 0.5 // bias
		}

		output := n.Activate(input)
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
