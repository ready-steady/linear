package lapack

// #cgo LDFLAGS: -lgfortran -lm
// #include "blas.h"
import "C"

// DGEMV performs one of the matrix-vector operations
//
//     y := alpha * A * x + beta * y or
//     y := alpha * A**T * x + beta * y
//
// where alpha and beta are scalars, x and y are vectors, and A is an m-by-n
// matrix.
//
// http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html
func DGEMV(trans byte, m, n int, alpha float64, A []float64, ldA int,
	x []float64, incx int, beta float64, y []float64, incy int) {

	C.dgemv(C.char(trans), C.int(m), C.int(n), C.double(alpha), (*C.double)(&A[0]), C.int(ldA),
		(*C.double)(&x[0]), C.int(incx), C.double(beta), (*C.double)(&y[0]), C.int(incy))
}

// DGEMM performs one of the matrix-matrix operations
//
//     C := alpha * op(A) * op(B) + beta * C
//
// where op(X) is one of
//
//     op(X) = X or
//     op(X) = X**T,
//
// alpha and beta are scalars, and A, B, and C are matrices, with op(A)
// an m-by-k matrix, op(B) a k-by-n matrix, and C an m-by-n matrix.
//
// http://www.netlib.org/lapack/explore-html/d7/d2b/dgemm_8f.html
func DGEMM(transa, transb byte, m, n, k int, alpha float64, A []float64,
	ldA int, B []float64, ldB int, beta float64, C []float64, ldC int) {

	C.dgemm(C.char(transa), C.char(transb), C.int(m), C.int(n), C.int(k), C.double(alpha), (*C.double)(&A[0]),
		C.int(ldA), (*C.double)(&B[0]), C.int(ldB), C.double(beta), (*C.double)(&C[0]), C.int(ldC))
}
