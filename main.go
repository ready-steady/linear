// Package linear provides a linear-algebra toolbox.
package linear

// TensorFloat64 computes the tensor product of a number of vectors.
func TensorFloat64(vectors ...[]float64) []float64 {
	nd := len(vectors)

	dimensions := make([]int, nd)
	for i := 0; i < nd; i++ {
		dimensions[i] = len(vectors[i])
	}

	ascend, descend := prepareTensor(dimensions)
	np := dimensions[0] * descend[0]

	tensor := make([]float64, np*nd)
	for i := 0; i < nd; i++ {
		for j, z := 0, i; j < descend[i]; j++ {
			for k := 0; k < dimensions[i]; k++ {
				for l := 0; l < ascend[i]; l++ {
					tensor[z] = vectors[i][k]
					z += nd
				}
			}
		}
	}

	return tensor
}

// TensorUint64 computes the tensor product of a number of vectors.
func TensorUint64(vectors ...[]uint64) []uint64 {
	nd := len(vectors)

	dimensions := make([]int, nd)
	for i := 0; i < nd; i++ {
		dimensions[i] = len(vectors[i])
	}

	ascend, descend := prepareTensor(dimensions)
	np := dimensions[0] * descend[0]

	tensor := make([]uint64, np*nd)
	for i := 0; i < nd; i++ {
		for j, z := 0, i; j < descend[i]; j++ {
			for k := 0; k < dimensions[i]; k++ {
				for l := 0; l < ascend[i]; l++ {
					tensor[z] = vectors[i][k]
					z += nd
				}
			}
		}
	}

	return tensor
}

func prepareTensor(dimensions []int) ([]int, []int) {
	nd := len(dimensions)

	ascend := make([]int, nd)
	ascend[0] = 1
	for i := 1; i < nd; i++ {
		ascend[i] = dimensions[i-1] * ascend[i-1]
	}

	descend := make([]int, nd)
	descend[nd-1] = 1
	for i := nd - 2; i >= 0; i-- {
		descend[i] = dimensions[i+1] * descend[i+1]
	}

	return ascend, descend
}
