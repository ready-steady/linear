package blas

import (
	"testing"
)

func TestDgemv(t *testing.T) {
	m := 2
	n := 3

	a := []float64{1, 4, 2, 5, 3, 6}
	x := []float64{1, 2, 3}
	y := []float64{6, 8}

	Dgemv('n', m, n, 1, a, m, x, 1, 1, y, 1)

	z := []float64{20, 40}

	for i := range y {
		if y[i] != z[i] {
			t.Error("The result is incorrect.")
		}
	}
}

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

func BenchmarkDgemm(t *testing.B) {
	m := 1000

	a := make([]float64, m * m)
	b := make([]float64, m * m)
	c := make([]float64, m * m)

	for i := 0; i < t.N; i++ {
		Dgemm('n', 'n', m, m, m, 1, a, m, b, m, 1, c, m)
	}
}
