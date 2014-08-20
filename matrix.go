package linear

import (
  "errors"
)

type Matrix struct {
  rows, cols uint
  data []float64
}

func (A *Matrix) Equal(B *Matrix) bool {
  if A.rows != B.rows || A.cols != B.rows {
    return false
  }

  for i := range A.data {
    if A.data[i] != B.data[i] {
      return false
    }
  }

  return true
}

func AddMatrix(A, B *Matrix) (Matrix, error) {
  if A.cols != B.rows {
    return Matrix{}, errors.New("A should have as many rows as B has columns.")
  }

  R := Matrix{
    rows: A.rows,
    cols: B.cols,
    data: make([]float64, A.rows * B.cols),
  }

  for i := range A.data {
    R.data[i] = A.data[i] + B.data[i]
  }

  return R, nil
}
