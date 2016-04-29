package matrix

import (
	"testing"

	"github.com/ready-steady/assert"
)

func BenchmarkMultiplyMatrixMatrix(bench *testing.B) {
	m := uint(1000)

	a := make([]float64, m*m)
	b := make([]float64, m*m)
	c := make([]float64, m*m)

	fillin(a, 1)
	fillin(b, 1)
	fillin(c, 1)

	for i := 0; i < bench.N; i++ {
		Multiply(a, b, c, m, m, m)
	}
}

func BenchmarkMultiplyMatrixVector(bench *testing.B) {
	m := uint(1000)

	a := make([]float64, m*m)
	b := make([]float64, m*1)
	c := make([]float64, m*m)

	fillin(a, 1)
	fillin(b, 1)
	fillin(c, 1)

	for i := 0; i < bench.N; i++ {
		Multiply(a, b, c, m, m, 1)
	}
}

func TestIdentity(t *testing.T) {
	assert.Equal(Identity(3), []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}, t)
}

func TestInverse(t *testing.T) {
	n := uint(3)
	a := []float64{1, -1, 0, 0, 5, 3, 2, 0, -9}

	expectedA := []float64{
		+8.823529411764706e-01, +1.764705882352941e-01, +5.882352941176472e-02,
		-1.176470588235294e-01, +1.764705882352941e-01, +5.882352941176472e-02,
		+1.960784313725490e-01, +3.921568627450981e-02, -9.803921568627452e-02,
	}

	assert.Equal(Inverse(a, n), nil, t)
	assert.EqualWithin(a, expectedA, 1e-15, t)
}

func TestMultiplyMatrixVector(t *testing.T) {
	m := uint(2)
	p := uint(4)
	n := uint(1)

	a := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	b := []float64{1, 2, 3, 4}
	c := make([]float64, m)

	Multiply(a, b, c, m, p, n)

	assert.Equal(c, []float64{50, 60}, t)
}

func TestMultiplyAdd(t *testing.T) {
	m := uint(2)
	p := uint(3)
	n := uint(4)

	a := []float64{1, 2, 3, 4, 5, 6}
	b := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	c := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	d := make([]float64, m*n)

	MultiplyAdd(a, b, c, d, m, p, n)

	assert.Equal(d, []float64{23, 30, 52, 68, 81, 106, 110, 144}, t)
	assert.Equal(c, []float64{1, 2, 3, 4, 5, 6, 7, 8}, t)

	MultiplyAdd(a, b, c, c, m, p, n)

	assert.Equal(c, []float64{23, 30, 52, 68, 81, 106, 110, 144}, t)
}

func fillin(a []float64, v float64) {
	for i := range a {
		a[i] = v
	}
}
