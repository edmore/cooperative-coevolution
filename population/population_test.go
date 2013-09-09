package population

import (
	"fmt"
	"testing"
)

func TestNewPopulation(t *testing.T) {
	p := NewPopulation(10)
	fmt.Println(p)
}

func TestCreate(t *testing.T) {
	p := NewPopulation(10)
	p.Create()
	fmt.Println(p)
}

func TestSelectNeuron(t *testing.T) {
	p := NewPopulation(10)
	p.Create()
	fmt.Println(p)
	fmt.Println(p.SelectNeuron())
}
