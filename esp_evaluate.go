package main

import (
	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
)

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network, c chan network.Network) {
	fitness := 0
	input := make([]float64, n.GetTotalInputs())
	output := make([]float64, n.GetTotalOutputs())

	for e.WithinTrackBounds() && e.WithinAngleBounds() && fitness < *goalFitness {
		state := e.GetState()
		input[0] = state.X / 4.8
		input[1] = state.XDot / 2
		input[2] = state.Theta1 / 0.52
		input[3] = state.Theta2 / 0.52
		input[4] = state.ThetaDot1 / 2
		input[5] = state.ThetaDot2 / 2

		if n.HasBias() {
			input[6] = 0.5 // bias
		}
		out := n.Activate(input, output)
		e.PerformAction(out[0])
		fitness++
	}
	// award fitness score to network
	n.SetFitness(fitness)
	c <- n
}
