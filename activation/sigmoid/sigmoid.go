package sigmoid

import "math"

// The Logistic (sigmoid) function

func Logistic(b float64, t float64) float64 {
	return (1 / (1 + math.Exp(-(b * t))))
}
