// Package matrix provides algorithms for manipulating real matrices.
package matrix

// Matrix represents a real matrix.
type Matrix struct {
	Rows, Cols uint32
	Data       []float64
}

// New creates a new matrix and sets its content to the given data without
// copying it.
func New(rows, cols uint32, data []float64) *Matrix {
	return &Matrix{rows, cols, data}
}

// Zero creates a new zeroed matrix.
func Zero(rows, cols uint32) *Matrix {
	return &Matrix{rows, cols, make([]float64, rows*cols)}
}

// Copy makes a copy of the matrix.
func (A *Matrix) Copy() *Matrix {
	B := &Matrix{A.Rows, A.Cols, make([]float64, A.Rows*A.Cols)}
	copy(B.Data, A.Data)

	return B
}

// IsEqual checks if two matrices are equal.
func (A *Matrix) IsEqual(B *Matrix) bool {
	if A.Rows != B.Rows || A.Cols != B.Cols {
		return false
	}

	for i := range A.Data {
		if B.Data[i] != B.Data[i] {
			return false
		}
	}

	return true
}

// IsSquare checks if the number of rows equals the number of columns.
func (A *Matrix) IsSquare() bool {
	return A.Rows == A.Cols
}
