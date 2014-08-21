// An interface for the Basic Linear Algebra Subprograms (BLAS)
//
// Reference: http://www.netlib.org/blas/
package blas

// #cgo LDFLAGS: -lblas
// #include "blas.h"
import "C"

import "unsafe"

// General matrix-vector multiplication according to the formula
//
// y := alpha * a * x + beta * y
//
// where
//
// a is an (m x n) matrix,
// x is an n-element vector, and
// y is an m-element vector.
//
// Reference: http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html
func Dgemv(trans byte, m, n int, alpha float64, a []float64, lda int,
	x []float64, incx int, beta float64, y []float64, incy int) {

	C.dgemv_(
		(*C.char)(unsafe.Pointer(&trans)),
		(*C.int)(unsafe.Pointer(&m)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.double)(unsafe.Pointer(&alpha)),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.int)(unsafe.Pointer(&lda)),
		(*C.double)(unsafe.Pointer(&x[0])),
		(*C.int)(unsafe.Pointer(&incx)),
		(*C.double)(unsafe.Pointer(&beta)),
		(*C.double)(unsafe.Pointer(&y[0])),
		(*C.int)(unsafe.Pointer(&incy)))
}

// General matrix-matrix multiplication according to the formula
//
// c := alpha * a * b + beta * c
//
// where
//
// a is an (m x k) matrix,
// b is a (k x n) matrix, and
// c is an (m x n) matrix.
//
// Reference: http://www.netlib.org/lapack/explore-html/d7/d2b/dgemm_8f.html
func Dgemm(transa, transb byte, m, n, k int, alpha float64, a []float64,
	lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {

	C.dgemm_(
		(*C.char)(unsafe.Pointer(&transa)),
		(*C.char)(unsafe.Pointer(&transb)),
		(*C.int)(unsafe.Pointer(&m)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.int)(unsafe.Pointer(&k)),
		(*C.double)(unsafe.Pointer(&alpha)),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.int)(unsafe.Pointer(&lda)),
		(*C.double)(unsafe.Pointer(&b[0])),
		(*C.int)(unsafe.Pointer(&ldb)),
		(*C.double)(unsafe.Pointer(&beta)),
		(*C.double)(unsafe.Pointer(&c[0])),
		(*C.int)(unsafe.Pointer(&ldc)))
}
