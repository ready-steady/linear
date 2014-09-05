// Package blas provides an interface to the Basic Linear Algebra Subprograms.
//
// http://www.netlib.org/blas/
package blas

// #cgo LDFLAGS: -Lbuild -lblas -lgfortran
// #include "blas.h"
import "C"

import "unsafe"

// Dgemv performs a general matrix-vector multiplication.
//
// The formula is as follows:
//
//     y := alpha * a * x + beta * y
//
// where a is an (m x n) matrix, x is an n-element vector, and y is an
// m-element vector.
//
// http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html
func Dgemv(trans byte, m, n int, alpha float64, a []float64, lda int,
	x []float64, incx int, beta float64, y []float64, incy int) {

	C.dgemv_(
		(*C.char)(unsafe.Pointer(&trans)),
		(*C.int)(unsafe.Pointer(&m)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.double)(&alpha),
		(*C.double)(&a[0]),
		(*C.int)(unsafe.Pointer(&lda)),
		(*C.double)(&x[0]),
		(*C.int)(unsafe.Pointer(&incx)),
		(*C.double)(&beta),
		(*C.double)(&y[0]),
		(*C.int)(unsafe.Pointer(&incy)))
}

// Dgemm performs a general matrix-matrix multiplication.
//
// The formula is as follows:
//
//     c := alpha * a * b + beta * c
//
// where a is an (m x k) matrix, b is a (k x n) matrix, and c is an (m x n)
// matrix.
//
// http://www.netlib.org/lapack/explore-html/d7/d2b/dgemm_8f.html
func Dgemm(transa, transb byte, m, n, k int, alpha float64, a []float64,
	lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {

	C.dgemm_(
		(*C.char)(unsafe.Pointer(&transa)),
		(*C.char)(unsafe.Pointer(&transb)),
		(*C.int)(unsafe.Pointer(&m)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.int)(unsafe.Pointer(&k)),
		(*C.double)(&alpha),
		(*C.double)(&a[0]),
		(*C.int)(unsafe.Pointer(&lda)),
		(*C.double)(&b[0]),
		(*C.int)(unsafe.Pointer(&ldb)),
		(*C.double)(&beta),
		(*C.double)(&c[0]),
		(*C.int)(unsafe.Pointer(&ldc)))
}
