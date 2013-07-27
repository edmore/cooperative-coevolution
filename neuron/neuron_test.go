package neuron

import "testing"

func TestNewNeuron(t *testing.T) {
	const in, out = 5, 5
	var n = NewNeuron(in)
	if x := len(n.weight); x != out {
		t.Errorf("len(NewNeuron(%v).weight) = %v, we want %v", in, x, out)
	}
}
