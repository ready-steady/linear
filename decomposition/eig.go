package decomposition

import (
	"fmt"

	"github.com/ready-steady/linear/lapack"
)

// SymEig performs the eigendecomposition of a symmetric m-by-m matrix A and
// stores the resulting eigenvectors and eigenvalues in an m-by-m matrix U and
// m-by-1 matrix Λ, respectively.
//
// https://en.wikipedia.org/wiki/Eigendecomposition_of_a_matrix#Real_symmetric_matrices
func SymEig(A, U, Λ []float64, m uint) error {
	if &A[0] != &U[0] {
		// NOTE: Only the upper triangular matrix is actually needed; however,
		// copying only that part might not be optimal for performance. Check!
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
