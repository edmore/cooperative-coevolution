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

func TestSetFitness(t *testing.T) {
	const out = 2
	n := NewNeuron(5)
	n.SetFitness(1)
	n.SetFitness(1)
	fmt.Println(n)
	if x := n.Fitness; x != out {
		t.Errorf("Fitness = %v, we want %v", x, out)
	}

}

func TestPerturb(t *testing.T) {
}
