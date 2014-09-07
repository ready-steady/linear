package matrix

import (
	"testing"

	"github.com/go-math/support/assert"
)

func TestAdd(t *testing.T) {
	m := uint32(3)
	n := uint32(3)

	A := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	B := []float64{0, 1, 2, 3, 4, 3, 2, 1, 0}
	C := make([]float64, m*n)

	Add(A, B, C, m, n)

	assert.Equal(C, []float64{1, 3, 5, 7, 9, 9, 9, 9, 9}, t)
}

func BenchmarkMultiplyMatrixMatrix(b *testing.B) {
	m := uint32(1000 + 1)
	p := uint32(1000 + 2)
	n := uint32(1000 + 3)

	A := make([]float64, m*p)
	B := make([]float64, p*n)
	C := make([]float64, m*n)

	for i := 0; i < b.N; i++ {
		Multiply(A, B, C, m, p, n)
	}
}

func BenchmarkMultiplyMatrixVector(b *testing.B) {
	m := uint32(1000)

	A := make([]float64, m*m)
	B := make([]float64, m*1)
	C := make([]float64, m*m)

	for i := 0; i < b.N; i++ {
		Multiply(A, B, C, m, m, 1)
	}
}
