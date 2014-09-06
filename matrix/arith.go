package matrix

import (
	"errors"

	"github.com/go-math/linal/lapack"
)

// Add performs matrix summation and returns the result in a new matrix.
func (A *Matrix) Add(B *Matrix) (*Matrix, error) {
	m, n := A.Rows, B.Cols

	if m != B.Rows || n != B.Cols {
		return nil, errors.New("the matrix dimensions are incompatible")
	}

	C := Zero(m, n)

	for i := range A.Data {
		B.Data[i] = A.Data[i] + B.Data[i]
	}

	return C, nil
}

// Multiply performs matrix multiplication and returns the result in a new
// matrix.
func (A *Matrix) Multiply(B *Matrix) (*Matrix, error) {
	m, n, k := A.Rows, B.Cols, A.Cols

	if k != B.Rows {
		return nil, errors.New("the matrix dimensions are incompatible")
	}

	C := Zero(m, n)

	if n == 1 {
		lapack.DGEMV('N', int(m), int(n), 1, A.Data, int(m), B.Data, 1, 0, C.Data, 1)
	} else {
		lapack.DGEMM('N', 'N', int(m), int(n), int(k), 1, A.Data, int(m), B.Data, int(k), 0, C.Data, int(m))
	}

	return C, nil
}
