package main

import (
	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"
)

var ch = make(chan network.Network)

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network) {
	fitness := 0
	// outputs := make([]float64, 0)
	for e.WithinTrackBounds() && e.WithinAngleBounds() {
		state := e.GetState()
		input := make([]float64, n.GetTotalInputs())
		input[0] = state.X / 4.8
		input[1] = state.XDot / 2
		input[2] = state.Theta1 / 0.52
		input[3] = state.Theta2 / 0.52
		input[4] = state.ThetaDot1 / 2
		input[5] = state.ThetaDot2 / 2

		if n.HasBias() {
			input[6] = 0.5 // bias
		}
		output := n.Activate(input)
		//	outputs = append(outputs, output[0])
		e.PerformAction(output[0])
		fitness++
	}
	// award fitness score to network
	n.SetFitness(fitness)
	// add the fitness score to cumulative fitness of
	// neurons that participated in trial.
	// beware of race conditions when adding the fitness
	// to each neuron needs to be synchronized.
	//	fmt.Println(outputs)
	ch <- n
}
