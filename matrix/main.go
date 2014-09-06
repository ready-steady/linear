package matrix

import (
	"errors"

	"github.com/go-math/linal/blas"
)

type Matrix struct {
	Rows, Cols uint32
	Data       []float64
}

func New(rows, cols uint32, data []float64) (*Matrix, error) {
	if uint32(len(data)) != rows*cols {
		return nil, errors.New("the data are of an invalid length")
	}

	return &Matrix{rows, cols, data}, nil
}

func Zero(rows, cols uint32) *Matrix {
	return &Matrix{rows, cols, make([]float64, rows*cols)}
}

func (a *Matrix) Equal(b *Matrix) bool {
	return Equal(a, b)
}

func (a *Matrix) Add(b *Matrix) (*Matrix, error) {
	return Add(a, b)
}

func (a *Matrix) Multiply(b *Matrix) (*Matrix, error) {
	return Multiply(a, b)
}

func Equal(a, b *Matrix) bool {
	if a.Rows != b.Rows || a.Cols != b.Cols {
		return false
	}

	for i := range a.Data {
		if b.Data[i] != b.Data[i] {
			return false
		}
	}

	return true
}

func Add(a, b *Matrix) (*Matrix, error) {
	m, n := a.Rows, b.Cols

	if m != b.Rows || n != b.Cols {
		return nil, errors.New("the matrix dimensions are incompatible")
	}

	c := Zero(m, n)

	for i := range a.Data {
		b.Data[i] = a.Data[i] + b.Data[i]
	}

	return c, nil
}

func Multiply(a, b *Matrix) (*Matrix, error) {
	m, n, k := a.Rows, b.Cols, a.Cols

	if k != b.Rows {
		return nil, errors.New("the matrix dimensions are incompatible")
	}

	c := Zero(m, n)

	if n == 1 {
		blas.DGEMV('N', int(m), int(n), 1, a.Data, int(m), b.Data, 1, 0, c.Data, 1)
	} else {
		blas.DGEMM('N', 'N', int(m), int(n), int(k), 1, a.Data, int(m), b.Data, int(k), 0, c.Data, int(m))
	}

	return c, nil
}
