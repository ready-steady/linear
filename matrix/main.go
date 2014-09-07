// Package matrix provides algorithms for manipulating real matrices.
package matrix

import (
	"github.com/go-math/linal/lapack"
)

// Add performs summation of an m-by-n matrix A with an m-by-n matrix B and
// stores the result in a m-by-n matrix C.
func Add(A, B, C []float64, m, n uint32) {
	for i := range C {
		C[i] = A[i] + B[i]
	}
}

// Multiply performs multiplication of an m-by-p matrix A with a p-by-n matrix
// B and stores the result in an m-by-n matrix C.
func Multiply(A, B, C []float64, m, p, n uint32) {
	if n == 1 {
		lapack.DGEMV('N', int(m), int(n), 1, A, int(m), B, 1, 0, C, 1)
	} else {
		lapack.DGEMM('N', 'N', int(m), int(n), int(p), 1, A, int(m), B, int(p), 0, C, int(m))
	}
}
