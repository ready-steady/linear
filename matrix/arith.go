package matrix

import (
	"errors"

	"github.com/go-math/linal/blas"
)

func (a *Matrix) Add(b *Matrix) (*Matrix, error) {
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

func (a *Matrix) Multiply(b *Matrix) (*Matrix, error) {
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
