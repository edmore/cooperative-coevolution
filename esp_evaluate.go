package main

import (
	"encoding/json"
	"fmt"

	"github.com/edmore/esp/environment"
	"github.com/edmore/esp/network"

	"io/ioutil"
)

// Evaluate the network in the trial environment
func evaluate(e environment.Environment, n network.Network, c chan network.Network) {
	fitness := 0
	input := make([]float64, n.GetTotalInputs())
	output := make([]float64, n.GetTotalOutputs())
	var state *environment.State
	states := make([]environment.State, 0)

	for e.WithinTrackBounds() && e.WithinAngleBounds() && fitness < *goalFitness {
		state = e.GetState()
		// push state into states slice
		states = append(states, *state)
		// Proceed to next state
		if *markov == true {
			input[0] = state.X / 4.8
			input[1] = state.XDot / 2
			input[2] = state.Theta1 / 0.52
			input[3] = state.Theta2 / 0.52
			input[4] = state.ThetaDot1 / 2
			input[5] = state.ThetaDot2 / 2
			if n.HasBias() {
				input[6] = 0.5 // bias
			}
		} else {
			input[0] = state.X / 4.8
			input[1] = state.Theta1 / 0.52
			input[2] = state.Theta2 / 0.52
			if n.HasBias() {
				input[3] = 0.5 // bias
			}
		}

		out := n.Activate(input, output)
		e.PerformAction(out[0])
		fitness++
	}

	if *simulation == true {
		if fitness == *goalFitness {
			// write the states to a json file
			b, err := json.Marshal(states)
			if err != nil {
				fmt.Println("error:", err)
			}
			err = ioutil.WriteFile("simulation/processingjs/json/states.json", b, 0644)
			if err != nil {
				panic(err)
			}
		}
	}
	// award fitness score to network
	n.SetFitness(fitness)
	c <- n
}
