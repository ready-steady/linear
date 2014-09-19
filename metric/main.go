// Package metric provides functions for computing metrics.
package metric

import (
	"math"
)

// L2 computes the Euclidean distance between two n-element vectors in the
// n-dimensional Euclidean space.
func L2(x, y []float64) float64 {
	var sum, Δ float64

	for i := range x {
		Δ = x[i] - y[i]
		sum += Δ * Δ
	}

	return math.Sqrt(sum)
}
