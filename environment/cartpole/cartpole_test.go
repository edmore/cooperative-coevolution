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

	fmt.Println("Initial State:", c.GetState())
	for c.WithinTrackBounds() && c.WithinAngleBounds() {
		s := c.PerformAction(1)
		fmt.Println("New State", s)
	}
}
