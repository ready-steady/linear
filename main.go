// Package linear provides a linear-algebra toolbox.
package linear

func Tensor(data ...[]float64) []float64 {
	nd := len(data)

	dims := make([]int, nd)
	for i := 0; i < nd; i++ {
		dims[i] = len(data[i])
	}

	aprod := make([]int, nd)
	aprod[0] = 1
	for i := 1; i < nd; i++ {
		aprod[i] = dims[i-1] * aprod[i-1]
	}

	dprod := make([]int, nd)
	dprod[nd-1] = 1
	for i := nd - 2; i >= 0; i-- {
		dprod[i] = dims[i+1] * dprod[i+1]
	}

	np := dims[0] * dprod[0]

	tensor := make([]float64, np*nd)
	for i, z := 0, 0; i < nd; i++ {
		for j := 0; j < dprod[i]; j++ {
			for k := 0; k < dims[i]; k++ {
				for l := 0; l < aprod[i]; l++ {
					tensor[z] = data[i][k]
					z++
				}
			}
		}
	}

	return tensor
}
