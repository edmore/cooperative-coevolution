package neuron

import (
	"fmt"
	"testing"
)

func TestNewNeuron(t *testing.T) {
	const in, out = 5, 5
	n := NewNeuron(in)
	fmt.Println(n)
	if x := len(n.Weight); x != out {
		t.Errorf("len(NewNeuron(%v).Weight) = %v, we want %v", in, x, out)
	}
}

func TestCreate(t *testing.T) {
	n := NewNeuron(5)
	n.Create()
	fmt.Println(n)
	for _, w := range n.Weight {
		if w == 0 {
			t.Errorf("This value should not be zero")
		}
	}
}

func TestPerturb(t *testing.T) {
	n := NewNeuron(5)
	n.Create()
	fmt.Println(n)
	old_Weight := n.Weight[1]

	n.Perturb()
	fmt.Println(n)
	new_Weight := n.Weight[1]

	if old_Weight == new_Weight {
		t.Errorf("The Weights should be perturbed")
	}
}
