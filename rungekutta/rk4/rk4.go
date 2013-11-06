package rk4

/*
 Package rk4 implements the Runge-Kutta Fourth Order (Rk4) Method
*/

/*********************************

dx - step size / change in x
dydx - the differential equation

**********************************/

type Point struct {
	X, Y float64
}

type Equation func(float64, float64) float64

// Initialise a new point
func NewPoint(x float64, y float64) *Point {
	return &Point{X: x, Y: y}
}

// Step
func step(dx float64, dydx Equation, p *Point) {
	k1 := dydx(p.X, p.Y)
	k2 := dydx(p.X+(0.5*dx), p.Y+(0.5*k1*dx))
	k3 := dydx(p.X+(0.5*dx), p.Y+(0.5*k2*dx))
	k4 := dydx(p.X+dx, p.Y+k3*dx)
	// Update points
	p.X += dx
	p.Y = p.Y + (1.0/6.0)*(k1+2*k2+2*k3+k4)*dx
}

// Solve for y = f(x)
func (p *Point) Solve(dx float64, dydx Equation, xFinal float64) float64 {
	for p.X < xFinal {
		step(dx, dydx, p)
	}
	return p.Y
}
