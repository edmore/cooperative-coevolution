/*
Package neuron implements an Artificial Neural Network neuron.
*/

package neuron

type Neuron struct {
	weight   []float64
	lesioned bool
	trials   int
	fitness  float64
}

// Neuron constructor
func NewNeuron(size int) *Neuron {
	var n *Neuron = new(Neuron)
	n.weight = make([]float64, size)
	return n
}

func (n *Neuron) Create() {
}
