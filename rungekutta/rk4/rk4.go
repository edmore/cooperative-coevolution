package rk4

/*
 Package classic implements the Runge-Kutta Fourth Order (Rk4) Method
*/

/*********************************

dx - step size / change in x
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
	temp := new(Point)
	temp.X, temp.Y = p.X, p.Y

	k1 := dydx(temp)
	temp.X, temp.Y = p.X+(0.5*dx), p.Y+(0.5*k1*dx)
	k2 := dydx(temp)
	temp.X, temp.Y = p.X+(0.5*dx), p.Y+(0.5*k2*dx)
	k3 := dydx(temp)
	temp.X, temp.Y = p.X+dx, p.Y+k3*dx
	k4 := dydx(temp)
	p.Y = p.Y + (1.0/6.0)*(k1+2*k2+2*k3+k4)*dx
	p.X = temp.X
}

// Solve for y = f(x)
func (p *Point) Solve(dx float64, dydx Equation, xFinal float64) float64 {
	for p.X < xFinal {
		step(dx, dydx, p)
	}
	return p.Y
}
