package neuron

import (
	"fmt"
	"testing"
)

func TestNewNeuron(t *testing.T) {
	const in, out = 5, 5
	var n = NewNeuron(in)
	if x := len(n.weight); x != out {
		t.Errorf("len(NewNeuron(%v).weight) = %v, we want %v", in, x, out)
	}
}

func TestCreate(t *testing.T) {
	var n = NewNeuron(5)
	n.Create()
	fmt.Println(n)
	for _, w := range n.weight {
		if w == 0 {
			t.Errorf("This value should not be zero")
		}
	}
}
