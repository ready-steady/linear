package decomp

import (
	"errors"
	"fmt"

	"github.com/go-math/linal/lapack"
	"github.com/go-math/linal/matrix"
)

// SymEig performs the eigendecomposition assuming that the matrix is
// symmetric:
//
//     A = Q * diag(lambda) * transpose(Q)
//
// where
//
//     A is a symmetric m-by-m matrix,
//     Q is an orthogonal m-by-m matrix of the eigenvectors of A, and
//     lambda is an m-element vector of the eigenvalues of A.
//
// https://en.wikipedia.org/wiki/Eigendecomposition_of_a_matrix#Real_symmetric_matrices
func SymEig(A *matrix.Matrix) (Q *matrix.Matrix, lambda *matrix.Matrix, err error) {
	if !A.IsSquare() {
		return nil, nil, errors.New("the matrix should be square")
	}

	m := A.Rows

	// TODO: Only the upper triangular matrix is actually needed; however,
	// copying only that part might not be optimal for performance. Check!
	Q = A.Copy()

	lambda = matrix.Zero(m, 1)

	// The size of the temporary array should have at least 3*m-1.
	temp := make([]float64, 4*m)
	flag := 0

	lapack.DSYEV('N', 'U', int(m), Q.Data, int(m), lambda.Data, temp, len(temp), &flag)

	if flag != 0 {
		return nil, nil, fmt.Errorf("LAPACK failed with code %v", flag)
	}

	return Q, lambda, nil
}
