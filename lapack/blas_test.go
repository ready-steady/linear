package lapack

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestDGEMV(t *testing.T) {
	M, N := 2, 3

	A := []float64{1, 4, 2, 5, 3, 6}
	X := []float64{1, 2, 3}
	Y := []float64{6, 8}

	DGEMV('N', M, N, 1, A, M, X, 1, 1, Y, 1)

	assert.Equal(Y, []float64{20, 40}, t)
}

func TestDGEMM(t *testing.T) {
	M, N, K := 2, 4, 3

	A := []float64{1, 4, 2, 5, 3, 6}
	B := []float64{1, 5, 9, 2, 6, 10, 3, 7, 11, 4, 8, 12}
	C := []float64{2, 7, 6, 2, 0, 7, 4, 2}

	DGEMM('N', 'N', M, N, K, 1, A, M, B, K, 1, C, M)

	assert.Equal(C, []float64{40, 90, 50, 100, 50, 120, 60, 130}, t)
}

func BenchmarkDGEMVFewLarge(b *testing.B) {
	M := 1000
	A, X, Y := ones(M*M), ones(M*1), ones(M*1)

	for i := 0; i < b.N; i++ {
		DGEMV('N', M, M, 1, A, M, X, 1, 1, Y, 1)
	}
}

func BenchmarkDGEMVManySmall(b *testing.B) {
	M := 20
	A, X, Y := ones(M*M), ones(M*1), ones(M*1)

	for i := 0; i < b.N; i++ {
		for j := 0; j < 20000; j++ {
			DGEMV('N', M, M, 1, A, M, X, 1, 1, Y, 1)
		}
	}
}

func BenchmarkDGEMM(b *testing.B) {
	M := 1000
	A, B, C := ones(M*M), ones(M*M), ones(M*M)

	for i := 0; i < b.N; i++ {
		DGEMM('N', 'N', M, M, M, 1, A, M, B, M, 1, C, M)
	}
}

func ones(size int) []float64 {
	data := make([]float64, size)
	for i := range data {
		data[i] = 1
	}
	return data
}
