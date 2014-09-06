package matrix

import (
	"errors"
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
