package matrix

import (
  "errors"
)

type Matrix struct {
  rows, cols int
  data []float64
}

func New(rows, cols int, data []float64) (Matrix, error) {
  switch len(data) {
  case 0:
    return Matrix{rows, cols, make([]float64, rows * cols)}, nil
  case rows * cols:
    return Matrix{rows, cols, data}, nil
  default:
    return Matrix{}, errors.New("The data are of an invalid length.")
  }
}

func (A *Matrix) Equal(B *Matrix) bool {
  return MatrixEqual(A, B)
}

func (A *Matrix) Add(B *Matrix) (Matrix, error) {
  return MatrixAdd(A, B)
}

func (A *Matrix) Multiply(B *Matrix) (Matrix, error) {
  return MatrixMultiply(A, B)
}

func MatrixEqual(A, B *Matrix) bool {
  if A.rows != B.rows || A.cols != B.cols {
    return false
  }

  for i := range A.data {
    if A.data[i] != B.data[i] {
      return false
    }
  }

  return true
}

func MatrixAdd(A, B *Matrix) (Matrix, error) {
  m, n := A.rows, A.cols

  if m != B.rows || n != B.cols {
    return Matrix{}, errors.New("A should have the same dimensions as B.")
  }

  R := Matrix{ rows: m, cols: n, data: make([]float64, m * n) }

  for i := range A.data {
    R.data[i] = A.data[i] + B.data[i]
  }

  return R, nil
}

func MatrixMultiply(A, B *Matrix) (Matrix, error) {
  m, n, p := A.rows, A.cols, B.cols

  if n != B.rows {
    return Matrix{}, errors.New("A should have as many columns as B has rows.")
  }

  R := Matrix{ rows: m, cols: p, data: make([]float64, m * p) }

  var sum float64
  for i := 0; i < m; i++ {
    for j := 0; j < p; j++ {
      sum = 0
      for k := 0; k < n; k++ {
        sum += A.data[i * n + k] * B.data[k * p + j]
      }
      R.data[i * p + j] = sum
    }
  }

  return R, nil
}
