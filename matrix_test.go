package linear

import "testing"

func TestAddMatrix(t *testing.T) {
  A := Matrix{ 3, 3, []float64{0, 1, 2, 3, 4, 3, 2, 1, 0} }
  B := Matrix{ 3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9} }
  C := Matrix{ 3, 3, []float64{1, 3, 5, 7, 9, 9, 9, 9, 9} }

  R, err := AddMatrix(&A, &B)

  if err != nil {
    t.Error("There should be no error.")
  }

  if !R.Equal(&C) {
    t.Error("The equality 'A + B = C' should hold.")
  }
}
