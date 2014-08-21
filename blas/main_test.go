package blas

import (
  "testing"
)

func TestDgemm(t *testing.T) {
  m := 2
  n := 4
  k := 3

  a := []float64{1, 4, 2, 5, 3, 6}
  b := []float64{1, 5, 9, 2, 6, 10, 3, 7, 11, 4, 8, 12}
  c := []float64{2, 7, 6, 2, 0, 7, 4, 2}

  Dgemm('n', 'n', m, n, k, 1, a, m, b, k, 1, c, m)

  d := []float64{40, 90, 50, 100, 50, 120, 60, 130}

  for i := range c {
    if c[i] != d[i] {
      t.Error("The result is incorrect.")
    }
  }
}
