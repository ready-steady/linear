// Package matrix provides algorithms for manipulating matrices.
package matrix

import (
	"github.com/ready-steady/lapack"
)

// Multiply multiplies an m-by-p matrix A by a p-by-n matrix B and stores the
// result in an m-by-n matrix C.
func Multiply(A, B, C []float64, m, p, n uint) {
	if n == 1 {
		lapack.DGEMV('N', int(m), int(p), 1, A, int(m), B, 1, 0, C, 1)
	} else {
		lapack.DGEMM('N', 'N', int(m), int(n), int(p), 1, A, int(m), B, int(p), 0, C, int(m))
	}
}

// MultiplyAdd multiplies an m-by-p matrix A by a p-by-n matrix B, sums the
// resulting m-by-n matrix with an m-by-n matrix C, and stores the final result
// in an m-by-n matrix D.
func MultiplyAdd(A, B, C, D []float64, m, p, n uint) {
	if &C[0] != &D[0] {
		copy(D, C)
	}

	if n == 1 {
		lapack.DGEMV('N', int(m), int(p), 1, A, int(m), B, 1, 1, D, 1)
	} else {
		lapack.DGEMM('N', 'N', int(m), int(n), int(p), 1, A, int(m), B, int(p), 1, D, int(m))
	}
}
