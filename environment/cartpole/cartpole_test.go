package cartpole

import (
	"fmt"
	"testing"
)

func TestNewCartpole(t *testing.T) {
	//c := NewCartpole()
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
	for state := c.GetState(); state.X > -2.4 && state.X < 2.4; {
		s := c.PerformAction(0)
		fmt.Println("New State", s)
	}
}
