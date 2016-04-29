// Package matrix provides basic matrix operations.
package matrix

import (
	"github.com/ready-steady/lapack"
)

// Identity constructs an m-by-m identity matrix.
func Identity(m uint) []float64 {
	a := make([]float64, m*m)
	for i := uint(0); i < m; i++ {
		a[i*m+i] = 1
	}
	return a
}

// Multiply multiplies an m-by-p matrix A by a p-by-n matrix B and stores the
// result in an m-by-n matrix C.
func Multiply(a, b, c []float64, m, p, n uint) {
	if n == 1 {
		lapack.DGEMV('N', int(m), int(p), 1, a, int(m), b, 1, 0, c, 1)
	} else {
		lapack.DGEMM('N', 'N', int(m), int(n), int(p), 1, a, int(m), b, int(p), 0, c, int(m))
	}
}

// MultiplyAdd multiplies an m-by-p matrix A by a p-by-n matrix B, sums the
// resulting m-by-n matrix with an m-by-n matrix C, and stores the final result
// in an m-by-n matrix D.
func MultiplyAdd(a, b, c, d []float64, m, p, n uint) {
	if &c[0] != &d[0] {
		copy(d, c)
	}

	if n == 1 {
		lapack.DGEMV('N', int(m), int(p), 1, a, int(m), b, 1, 1, d, 1)
	} else {
		lapack.DGEMM('N', 'N', int(m), int(n), int(p), 1, a, int(m), b, int(p), 1, d, int(m))
	}
}
