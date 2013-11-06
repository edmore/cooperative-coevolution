package euler

import "testing"

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
	eq := func(x, y float64) float64 { return 2 * (x - 1) }
	const x, y, dx, xFinal, yFinal = 1, 0, 0.5, 3, 3
	initialPoint := NewPoint(x, y)
	solution := initialPoint.Solve(dx, eq, xFinal)

	if solution != yFinal {
		t.Errorf("value of yFinal should be %v, we got %v", yFinal, solution)
	}

}
