package blas

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestDGEMV(t *testing.T) {
	m, n := 2, 3

	A := []float64{1, 4, 2, 5, 3, 6}
	x := []float64{1, 2, 3}
	y := []float64{6, 8}

	DGEMV('N', m, n, 1, A, m, x, 1, 1, y, 1)

	assert.Equal(y, []float64{20, 40}, t)
}

func TestDGEMM(t *testing.T) {
	m, n, k := 2, 4, 3

	A := []float64{1, 4, 2, 5, 3, 6}
	B := []float64{1, 5, 9, 2, 6, 10, 3, 7, 11, 4, 8, 12}
	C := []float64{2, 7, 6, 2, 0, 7, 4, 2}

	DGEMM('N', 'N', m, n, k, 1, A, m, B, k, 1, C, m)

	assert.Equal(C, []float64{40, 90, 50, 100, 50, 120, 60, 130}, t)
}

func BenchmarkDGEMM(b *testing.B) {
	m := 1000

	A := make([]float64, m*m)
	B := make([]float64, m*m)
	C := make([]float64, m*m)

	for i := 0; i < b.N; i++ {
		DGEMM('N', 'N', m, m, m, 1, A, m, B, m, 1, C, m)
	}
}
