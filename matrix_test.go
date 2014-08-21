package linear

import "testing"

func TestMatrixAdd(t *testing.T) {
  A := Matrix{ 3, 3, []float{0, 1, 2, 3, 4, 3, 2, 1, 0} }
  B := Matrix{ 3, 3, []float{1, 2, 3, 4, 5, 6, 7, 8, 9} }
  C := Matrix{ 3, 3, []float{1, 3, 5, 7, 9, 9, 9, 9, 9} }

  R, err := A.Add(&B)

  if err != nil {
    t.Error("There should be no error.")
  }

  if !R.Equal(&C) {
    t.Error("The addition should work.")
  }
}

func TestMatrixAddError(t *testing.T) {
  A := Matrix{ 3, 3, []float{0, 1, 2, 3, 4, 3, 2, 1, 0} }
  B := Matrix{ 2, 3, []float{1, 2, 3, 4, 5, 6} }

  R, err := A.Add(&B)

  if err == nil {
    t.Error("There should be an error.")
  }

  if R.rows != 0 || R.cols != 0 || len(R.data) > 0 {
    t.Error("The result should be empty.")
  }
}

func TestMatrixMultiply(t *testing.T) {
  A := Matrix{ 3, 4, []float{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} }
  B := Matrix{ 4, 2, []float{1, 2, 3, 4, 5, 6, 7, 8} }
  C := Matrix{ 3, 2, []float{50, 60, 114, 140, 178, 220} }

  R, err := A.Multiply(&B)

  if err != nil {
    t.Error("There should be no error.")
  }

  if !R.Equal(&C) {
    t.Error("The multiplication should work.")
  }
}
