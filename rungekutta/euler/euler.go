/*
 Package euler implements the Runge-Kutta First Order (Euler) Method
*/

package euler

/*********************************

dx - step size / change in x
dy - change in y
dydx - the differential equation

**********************************/

type Point struct {
	X, Y float64
}

type Equation func(*Point) float64

// Initialise a new point
func NewPoint(x float64, y float64) *Point {
	return &Point{X: x, Y: y}
}

// Step
func step(dx float64, dydx Equation, p *Point) {
	dy := dx * dydx(p)
	p.X, p.Y = p.X+dx, p.Y+dy
}

// Solve for y = f(x)
func (p *Point) Solve(dx float64, dydx Equation, xFinal float64) float64 {
	for p.X < xFinal {
		step(dx, dydx, p)
	}
	return p.Y
}
