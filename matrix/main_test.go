package matrix

import "testing"

func BenchmarkMatrixMultiply(t *testing.B) {
	m := 1000

	a := Zero(m, m)
	b := Zero(m, m)

	for i := 0; i < t.N; i++ {
		_, _ = a.Multiply(&b)
	}
}
