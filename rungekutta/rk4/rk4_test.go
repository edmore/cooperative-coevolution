package rk4

import (
	"math"
	"testing"
)

func TestNewPoint(t *testing.T) {
	const x, y = 1, 0
	p := NewPoint(x, y)
	if p.X != 1 {
		t.Errorf("value of x should be %v, we got %v", x, p.X)
	}
	if p.Y != 0 {
		t.Errorf("value of x should be %v, we got %v", y, p.Y)
	}
}

func TestSolve(t *testing.T) {
	eq := func(p *Point) float64 {
		return -2.2067 * math.Pow(10, -12) * (math.Pow(p.Y, 4) - 81*math.Pow(10, 8))
	}
	const x, y, dx, xFinal, yFinal = 0, 1200, 240, 480, 594.912631110278
	initialPoint := NewPoint(x, y)
	solution := initialPoint.Solve(dx, eq, xFinal)

	if solution != yFinal {
		t.Errorf("value of yFinal should be %v, we got %v", yFinal, solution)
	}

}
