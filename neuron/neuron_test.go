package neuron

import (
	"fmt"
	"testing"
)

func TestNewNeuron(t *testing.T) {
	const in, out = 5, 5
	n := NewNeuron(in)
	fmt.Println(n)
	if x := len(n.weight); x != out {
		t.Errorf("len(NewNeuron(%v).weight) = %v, we want %v", in, x, out)
	}
}

func TestCreate(t *testing.T) {
	n := NewNeuron(5)
	n.Create()
	fmt.Println(n)
	for _, w := range n.weight {
		if w == 0 {
			t.Errorf("This value should not be zero")
		}
	}
}

func TestPerturb(t *testing.T) {
	n := NewNeuron(5)
	n.Create()
	fmt.Println(n)
	old_weight := n.weight[1]

	n.Perturb()
	fmt.Println(n)
	new_weight := n.weight[1]

	if old_weight == new_weight {
		t.Errorf("The weights should be perturbed")
	}
}
