package population

import (
	"fmt"
	"testing"
)

func TestNewPopulation(t *testing.T) {
	p := NewPopulation(10, 5)
	fmt.Println(p)
}

func TestCreate(t *testing.T) {
	const psize = 10
	p := NewPopulation(psize, 5)
	p.Create()
	if p.NumIndividuals != psize {
		t.Errorf("Size of population = %v, we want %v", p.NumIndividuals, psize)
	}
	fmt.Println(p)
}

func TestSelectNeuron(t *testing.T) {
	p := NewPopulation(10, 5)
	p.Create()
	fmt.Println(p)
	fmt.Println(p.SelectNeuron())
}
