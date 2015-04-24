package decomposition

import (
	"fmt"

	"github.com/ready-steady/lapack"
)

// SymmetricEigen computes the eigendecomposition of a symmetric m-by-m matrix
// A. The eigenvectors and eigenvalues are stored in an m-by-m matrix U and
// m-by-1 matrix Λ, respectively, in the ascending order of the eigenvalues.
//
// https://en.wikipedia.org/wiki/Eigendecomposition_of_a_matrix#Real_symmetric_matrices
func SymmetricEigen(A, U, Λ []float64, m uint) error {
	if &A[0] != &U[0] {
		// Only the upper triangular matrix actually needs to be copied.
		copy(U, A)
	}

	flag := 0
	size := []float64{0}

	lapack.DSYEV('V', 'U', int(m), U, int(m), Λ, size, -1, &flag)
	if flag != 0 {
		return fmt.Errorf("LAPACK failed with code %v", flag)
	}

	work := make([]float64, int(size[0]))

	lapack.DSYEV('V', 'U', int(m), U, int(m), Λ, work, len(work), &flag)
	if flag != 0 {
		return fmt.Errorf("LAPACK failed with code %v", flag)
	}

	return nil
}
