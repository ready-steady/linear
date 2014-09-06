// Package blas provides an interface to the Basic Linear Algebra Subprograms.
//
// http://www.netlib.org/blas/
package blas

// #cgo LDFLAGS: -Lbuild -lblas -lgfortran
// #include "blas.h"
import "C"

import "unsafe"

// DGEMV performs one of the matrix-vector operations
//
//     y := alpha * A * x + beta * y or
//     y := alpha * transpose(A) * x + beta * y
//
// where alpha and beta are scalars, x and y are vectors, and A is an m-by-n
// matrix.
//
// http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html
func DGEMV(trans byte, m, n int, alpha float64, A []float64, ldA int,
	x []float64, incx int, beta float64, y []float64, incy int) {

	C.dgemv_(
		(*C.char)(unsafe.Pointer(&trans)),
		(*C.int)(unsafe.Pointer(&m)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.double)(&alpha),
		(*C.double)(&A[0]),
		(*C.int)(unsafe.Pointer(&ldA)),
		(*C.double)(&x[0]),
		(*C.int)(unsafe.Pointer(&incx)),
		(*C.double)(&beta),
		(*C.double)(&y[0]),
		(*C.int)(unsafe.Pointer(&incy)),
	)
}

// DGEMM performs one of the matrix-matrix operations
//
//     C := alpha * op(A) * op(B) + beta * C
//
//  where op(X) is one of
//
//     op(X) = X or
//     op(X) = transpose(X),
//
// alpha and beta are scalars, and A, B, and C are matrices, with op(A)
// an m-by-k matrix, op(B) a k-by-n matrix, and C an m-by-n matrix.
//
// http://www.netlib.org/lapack/explore-html/d7/d2b/dgemm_8f.html
func DGEMM(transa, transb byte, m, n, k int, alpha float64, A []float64,
	ldA int, B []float64, ldB int, beta float64, C []float64, ldC int) {

	C.dgemm_(
		(*C.char)(unsafe.Pointer(&transa)),
		(*C.char)(unsafe.Pointer(&transb)),
		(*C.int)(unsafe.Pointer(&m)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.int)(unsafe.Pointer(&k)),
		(*C.double)(&alpha),
		(*C.double)(&A[0]),
		(*C.int)(unsafe.Pointer(&ldA)),
		(*C.double)(&B[0]),
		(*C.int)(unsafe.Pointer(&ldB)),
		(*C.double)(&beta),
		(*C.double)(&C[0]),
		(*C.int)(unsafe.Pointer(&ldC)),
	)
}
