// Package matrix provides basic matrix operations.
package matrix

import (
	"fmt"

	"github.com/ready-steady/lapack"
)

// Dot computes the dot product of two m-element vectors.
func Dot(a, b []float64, m uint) float64 {
	return lapack.DDOT(int(m), a, 1, b, 1)
}

// Identity constructs an m-by-m identity matrix.
func Identity(m uint) []float64 {
	a := make([]float64, m*m)
	for i := uint(0); i < m; i++ {
		a[i*m+i] = 1.0
	}
	return a
}

// Invert computes the inverse of an m-by-m matrix.
func Invert(a []float64, m uint) error {
	ipiv := make([]int, m+1)
	info := 0

	lapack.DGETRF(int(m), int(m), a, int(m), ipiv, &info)
	if info != 0 {
		return fmt.Errorf("LAPACK failed with code %v", info)
	}

	lwork := m * m
	work := make([]float64, lwork)

	lapack.DGETRI(int(m), a, int(m), ipiv, work, int(lwork), &info)
	if info != 0 {
		return fmt.Errorf("LAPACK failed with code %v", info)
	}

	return nil
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
