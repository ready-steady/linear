package decomp

import (
	"fmt"

	"github.com/go-math/linal/lapack"
)

// SymEig performs the eigendecomposition of a symmetric m-by-m matrix A and
// stores the resulting eigenvectors and eigenvalues in an m-by-m matrix U and
// m-by-1 matrix Λ, respectively.
//
// https://en.wikipedia.org/wiki/Eigendecomposition_of_a_matrix#Real_symmetric_matrices
func SymEig(A, U, Λ []float64, m uint32) error {
	if &A[0] != &U[0] {
		// NOTE: Only the upper triangular matrix is actually needed; however,
		// copying only that part might not be optimal for performance. Check!
		copy(U, A)
	}

	// The size of the temporary array should have at least 3*m-1.
	temp := make([]float64, 4*m)
	flag := 0

	lapack.DSYEV('V', 'U', int(m), U, int(m), Λ, temp, len(temp), &flag)

	if flag != 0 {
		return fmt.Errorf("LAPACK failed with code %v", flag)
	}

	return nil
}
