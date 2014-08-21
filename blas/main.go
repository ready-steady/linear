package blas

// #cgo LDFLAGS: -lblas
// #include "blas.h"
import "C"

import "unsafe"

// General matrix-matrix multiplication according to the formula
//
// C = alpha * A * B + beta * C
//
// where
//
// A is an (m x k) matrix,
// B is a (k x n) matrix, and
// C is an (m x n) matrix.
//
// All matrices are stored column-wise. For further details, refer to
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
		(*C.double)(unsafe.Pointer(&alpha)),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.int)(unsafe.Pointer(&lda)),
		(*C.double)(unsafe.Pointer(&b[0])),
		(*C.int)(unsafe.Pointer(&ldb)),
		(*C.double)(unsafe.Pointer(&beta)),
		(*C.double)(unsafe.Pointer(&c[0])),
		(*C.int)(unsafe.Pointer(&ldc)))
}
