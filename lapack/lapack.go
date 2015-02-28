package lapack

// #cgo LDFLAGS: -lgfortran -lm
// #include "lapack.h"
import "C"

import "unsafe"

// DGTSV computes the solution to system of linear equations A * X = B for
// triangular matrices.
//
// http://www.netlib.org/lapack/explore-html/d1/db3/dgtsv_8f.html
func DGTSV(N, NRHS int, DL, D, DU, B []float64, LDB int, INFO *int) {
	C.dgtsv(C.int(N), C.int(NRHS), (*C.double)(&DL[0]), (*C.double)(&D[0]),
		(*C.double)(&DU[0]), (*C.double)(&B[0]), C.int(LDB),
		(*C.int)(unsafe.Pointer(INFO)))
}

// DSYEV computes the eigenvalues and, optionally, the left and/or right
// eigenvectors for symmetric matrices.
//
// http://www.netlib.org/lapack/explore-html/dd/d4c/dsyev_8f.html
func DSYEV(JOBZ, UPLO byte, N int, A []float64, LDA int, W, WORK []float64,
	LWORK int, INFO *int) {

	C.dsyev(C.char(JOBZ), C.char(UPLO), C.int(N), (*C.double)(&A[0]), C.int(LDA),
		(*C.double)(&W[0]), (*C.double)(&WORK[0]), C.int(LWORK),
		(*C.int)(unsafe.Pointer(INFO)))
}
