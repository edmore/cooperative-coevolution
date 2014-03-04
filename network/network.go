package network

import "github.com/edmore/esp/neuron"

type Network interface {
	Activate()
	GetHiddenUnits() []*neuron.Neuron
}
