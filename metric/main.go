// Package metric provides functions for computing metrics.
package metric

import (
	"math"
)

// L2 computes the Euclidean distance between two n-element vectors.
func L2(x, y []float64) float64 {
	var sum, Δ float64

	for i := range x {
		Δ = x[i] - y[i]
		sum += Δ * Δ
	}

	return math.Sqrt(sum)
}

// Uniform computes the uniform distance between two n-element vectors.
func Uniform(x, y []float64) float64 {
	var δ, Δ float64

	for i := range x {
		δ = x[i] - y[i]
		if δ < 0 {
			δ = -δ
		}
		if δ > Δ {
			Δ = δ
		}
	}

	return Δ
}
