package linear

import (
  "errors"
)

type float float64

type Matrix struct {
  rows, cols uint
  data []float
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
    return Matrix{}, errors.New("B should have the same dimensions as A.")
  }

  R := Matrix{ rows: m, cols: n, data: make([]float, m * n) }

  for i := range A.data {
    R.data[i] = A.data[i] + B.data[i]
  }

  return R, nil
}

func MatrixMultiply(A, B *Matrix) (Matrix, error) {
  m, n, p := A.rows, A.cols, B.cols

  if n != B.rows {
    return Matrix{}, errors.New("B should have as many rows as A has columns.")
  }

  R := Matrix{ rows: m, cols: p, data: make([]float, m * p) }

  var i, j, k uint
  var sum float

  for i = 0; i < m; i++ {
    for j = 0; j < p; j++ {
      sum = 0

      for k = 0; k < n; k++ {
        sum += A.data[i * n + k] * B.data[k * p + j]
      }

      R.data[i * p + j] = sum
    }
  }

  return R, nil
}
