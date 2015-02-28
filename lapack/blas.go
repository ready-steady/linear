package lapack

// #cgo LDFLAGS: -lgfortran -lm
// #include "blas.h"
import "C"

// DGEMV performs one of the matrix-vector operations
//
//     Y := ALPHA * A * X + BETA * Y or
//     Y := ALPHA * A**T * X + BETA * Y
//
// where ALPHA and BETA are scalars, X and Y are vectors, and A is an M-by-N
// matrix.
//
// http://www.netlib.org/lapack/explore-html/dc/da8/dgemv_8f.html
func DGEMV(TRANS byte, M, N int, ALPHA float64, A []float64, LDA int,
	X []float64, INCX int, BETA float64, Y []float64, INCY int) {

	C.dgemv(C.char(TRANS), C.int(M), C.int(N), C.double(ALPHA),
		(*C.double)(&A[0]), C.int(LDA), (*C.double)(&X[0]), C.int(INCX),
		C.double(BETA), (*C.double)(&Y[0]), C.int(INCY))
}

// DGEMM performs one of the matrix-matrix operations
//
//     C := ALPHA * op(A) * op(B) + BETA * C
//
// where op(X) is one of
//
//     op(X) = X or
//     op(X) = X**T,
//
// ALPHA and BETA are scalars, and A, B, and C are matrices, with op(A)
// an M-by-K matrix, op(B) a K-by-N matrix, and C an M-by-N matrix.
//
// http://www.netlib.org/lapack/explore-html/d7/d2b/dgemm_8f.html
func DGEMM(TRANSA, TRANSB byte, M, N, K int, ALPHA float64, A []float64,
	LDA int, B []float64, LDB int, BETA float64, C []float64, LDC int) {

	C.dgemm(C.char(TRANSA), C.char(TRANSB), C.int(M), C.int(N), C.int(K),
		C.double(ALPHA), (*C.double)(&A[0]), C.int(LDA), (*C.double)(&B[0]),
		C.int(LDB), C.double(BETA), (*C.double)(&C[0]), C.int(LDC))
}
