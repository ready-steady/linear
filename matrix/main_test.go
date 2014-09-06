package matrix

import (
	"testing"
)

func TestAdd(t *testing.T) {
	A, _ := New(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	B, _ := New(3, 3, []float64{0, 1, 2, 3, 4, 3, 2, 1, 0})

	C, err := A.Add(B)

	if err != nil {
		t.Fatal("there should be no error")
	}

	D, _ := New(3, 3, []float64{1, 3, 5, 7, 9, 9, 9, 9, 9})

	if !C.Equal(D) {
		t.Fatal("the result is incorrect")
	}
}

func BenchmarkMultiplyMatrix(b *testing.B) {
	m := uint32(1000)

	A := Zero(m, m)
	B := Zero(m, m)

	for i := 0; i < b.N; i++ {
		A.Multiply(B)
	}
}

func BenchmarkMultiplyVector(b *testing.B) {
	m := uint32(1000)

	A := Zero(m, m)
	x := Zero(m, 1)

	for i := 0; i < b.N; i++ {
		A.Multiply(x)
	}
}
