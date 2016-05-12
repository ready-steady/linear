// Package metric provides functions for computing metrics.
//
// https://en.wikipedia.org/wiki/Metric_(mathematics)
package metric

import (
	"math"
)

// L2 computes the Euclidean distance between two vectors.
func L2(x, y []float64) float64 {
	sum := 0.0
	for i := range x {
		Δ := x[i] - y[i]
		sum += Δ * Δ
	}
	return math.Sqrt(sum)
}

// Uniform computes the uniform distance between two vectors.
func Uniform(x, y []float64) float64 {
	Δ := 0.0
	for i := range x {
		δ := x[i] - y[i]
		if δ < 0.0 {
			δ = -δ
		}
		if δ > Δ {
			Δ = δ
		}
	}
	return Δ
}
