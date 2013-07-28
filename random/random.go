package random

import (
	"math"
	"math/rand"
)

// Generate a random number from a cauchy distribution centered on zero.
func Cauchy(wtrange float64) float64 {
	u, Cauchy_cut := 0.5, 10.0

	for u == 0.5 {
		u = rand.Float64()
	}

	u = wtrange * math.Tan(u*math.Pi)
	if math.Abs(u) > Cauchy_cut {
		return Cauchy(wtrange)
	}
	return u
}
