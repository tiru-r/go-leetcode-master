package rotate_image_48

import "slices"

func rotate(matrix [][]int) {
	// transpose the matrix in-place using modern range-over-int
	for i := range len(matrix) {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each row using modern slices.Reverse with range-over-int
	for i := range len(matrix) {
		slices.Reverse(matrix[i])
	}
}
