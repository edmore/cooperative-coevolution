package cartpole

import (
	"fmt"
	"testing"
)

func TestNewCartpole(t *testing.T) {
	c := NewCartpole()
	fmt.Println(c)
}

func TestReset(t *testing.T) {
	c := NewCartpole()
	c.Reset()
	//	fmt.Println(s)
}

func TestPerformAction(t *testing.T) {
	c := NewCartpole()
	c.Reset()

	fmt.Println("Og State:", c.GetState())
	s := c.PerformAction(1)
	fmt.Println("New State", s)
}
