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
	s := c.Reset()
	fmt.Println(s)
}

func TestPerformAction(t *testing.T) {
}