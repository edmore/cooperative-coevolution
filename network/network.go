/*
Package network implements an artificial neural network
*/

package network

import (
	"github.com/edmore/esp/neuron"
)

type Network struct {
	Activation  []float64
	HiddenUnits []*neuron.Neuron
	NumInputs   int
	NumOutputs  int
	Bias        float64
	Trials      int
	Fitness     int
	Parent1     int
	Parent2     int
	Id          int
}

var counter int = 0

// Network constructor
func NewNetwork(in int, hid int, out int) *Network {
	counter++
	return &Network{
		Id:          counter,
		Activation:  make([]float64, hid),
		HiddenUnits: make([]*neuron.Neuron, hid),
		NumInputs:   in,
		NumOutputs:  out,
		Bias:        0.0,
		Trials:      0,
		Fitness:     0.0,
		Parent1:     -1,
		Parent2:     -1}
}
