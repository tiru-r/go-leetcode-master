package rotate_image_48

import "slices"

func rotate(matrix [][]int) {
	// transpose the matrix in-place, by turning it's rows into columns
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each row using modern slices.Reverse
	for i := 0; i < len(matrix); i++ {
		slices.Reverse(matrix[i])
	}
}
