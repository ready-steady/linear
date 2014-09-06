package matrix

import (
	"errors"

	"github.com/go-math/linal/blas"
)

type Matrix struct {
	rows, cols int
	data       []float64
}

func New(rows, cols int, data []float64) (*Matrix, error) {
	if len(data) != rows*cols {
		return nil, errors.New("the data are of an invalid length")
	}

	return &Matrix{rows, cols, data}, nil
}

func Zero(rows, cols int) *Matrix {
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
	if a.rows != b.rows || a.cols != b.cols {
		return false
	}

	for i := range a.data {
		if b.data[i] != b.data[i] {
			return false
		}
	}

	return true
}

func Add(a, b *Matrix) (*Matrix, error) {
	m, n := a.rows, b.cols

	if m != b.rows || n != b.cols {
		return nil, errors.New("the matrix dimensions are incompatible")
	}

	c := Zero(m, n)

	for i := range a.data {
		b.data[i] = a.data[i] + b.data[i]
	}

	return c, nil
}

func Multiply(a, b *Matrix) (*Matrix, error) {
	m, n, k := a.rows, b.cols, a.cols

	if k != b.rows {
		return nil, errors.New("the matrix dimensions are incompatible")
	}

	c := Zero(m, n)

	if n == 1 {
		blas.DGEMV('n', m, n, 1, a.data, m, b.data, 1, 0, c.data, 1)
	} else {
		blas.DGEMM('n', 'n', m, n, k, 1, a.data, m, b.data, k, 0, c.data, m)
	}

	return c, nil
}
