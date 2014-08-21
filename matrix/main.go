package matrix

import (
	"errors"
	"github.com/gomath/linear/blas"
)

type Matrix struct {
	rows, cols int
	data       []float64
}

func New(rows, cols int, data []float64) (Matrix, error) {
	if len(data) != rows*cols {
		return Matrix{}, errors.New("The data are of an invalid length.")
	}

	return Matrix{rows, cols, data}, nil
}

func Zero(rows, cols int) Matrix {
	return Matrix{rows, cols, make([]float64, rows*cols)}
}

func (a *Matrix) IsEqual(b *Matrix) bool {
	return IsEqual(a, b)
}

func (a *Matrix) Add(b *Matrix) (Matrix, error) {
	return Add(a, b)
}

func (a *Matrix) Multiply(b *Matrix) (Matrix, error) {
	return Multiply(a, b)
}

func IsEqual(a, b *Matrix) bool {
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

func Add(a, b *Matrix) (Matrix, error) {
	m, n := a.rows, b.cols

	if m != b.rows || n != b.cols {
		return Matrix{}, errors.New("The matrix dimensions are incompatible.")
	}

	c := Zero(m, n)

	for i := range a.data {
		b.data[i] = a.data[i] + b.data[i]
	}

	return c, nil
}

func Multiply(a, b *Matrix) (Matrix, error) {
	m, n, k := a.rows, b.cols, a.cols

	if k != b.rows {
		return Matrix{}, errors.New("The matrix dimensions are incompatible.")
	}

	c := Zero(m, n)

	blas.Dgemm('n', 'n', m, n, k, 1, a.data, m, b.data, k, 1, c.data, m)

	return c, nil
}
