// Package lapack provides an interface to the Linear Algebra PACKage.
//
// http://www.netlib.org/lapack/
package lapack

// #cgo LDFLAGS: -L. -llapack_gomath -lrefblas_gomath -lgfortran
// #include "lapack.h"
import "C"

import "unsafe"

// DSYEV computes the eigenvalues and, optionally, the left and/or right
// eigenvectors for symmetric matrices.
//
// http://www.netlib.org/lapack/explore-html/dd/d4c/dsyev_8f.html
func DSYEV(jobz, uplo byte, n int, A []float64, ldA int, w, work []float64,
	lwork int, info *int) {

	C.dsyev_(
		(*C.char)(unsafe.Pointer(&jobz)),
		(*C.char)(unsafe.Pointer(&uplo)),
		(*C.int)(unsafe.Pointer(&n)),
		(*C.double)(&A[0]),
		(*C.int)(unsafe.Pointer(&ldA)),
		(*C.double)(&w[0]),
		(*C.double)(&work[0]),
		(*C.int)(unsafe.Pointer(&lwork)),
		(*C.int)(unsafe.Pointer(info)))
}
