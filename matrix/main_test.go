package matrix

import "testing"

func TestAdd(t *testing.T) {
	a, _ := New(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	b, _ := New(3, 3, []float64{0, 1, 2, 3, 4, 3, 2, 1, 0})

	c, err := a.Add(&b)

	if err != nil {
		t.Error("There should be no error.")
	}

	d, _ := New(3, 3, []float64{1, 3, 5, 7, 9, 9, 9, 9, 9})

	if !c.IsEqual(&d) {
		t.Error("The result is incorrect.")
	}
}

func BenchmarkMultiply(t *testing.B) {
	m := 1000

	a := Zero(m, m)
	b := Zero(m, m)

	for i := 0; i < t.N; i++ {
		_, _ = a.Multiply(&b)
	}
}
