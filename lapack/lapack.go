package lapack

// #cgo LDFLAGS: -lgfortran -lm
// #include "lapack.h"
import "C"

import "unsafe"

// DSYEV computes the eigenvalues and, optionally, the left and/or right
// eigenvectors for symmetric matrices.
//
// http://www.netlib.org/lapack/explore-html/dd/d4c/dsyev_8f.html
func DSYEV(jobz, uplo byte, n int, A []float64, ldA int, w, work []float64,
	lwork int, info *int) {

	C.dsyev(C.char(jobz), C.char(uplo), C.int(n), (*C.double)(&A[0]), C.int(ldA),
		(*C.double)(&w[0]), (*C.double)(&work[0]), C.int(lwork), (*C.int)(unsafe.Pointer(info)))
}
