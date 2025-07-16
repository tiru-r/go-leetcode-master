package set_matrix_zeros_73

func setZeroes(matrix [][]int) {
	setColZero := false
	
	for r := range len(matrix) {
		for c := range len(matrix[r]) {
			if matrix[r][c] == 0 {
				if c == 0 {
					setColZero = true
				} else {
					matrix[r][0] = 0
					matrix[0][c] = 0
				}
			}
		}
	}
	
	for r := range len(matrix) - 1 {
		for c := range len(matrix[r+1]) - 1 {
			if matrix[r+1][0] == 0 || matrix[0][c+1] == 0 {
				matrix[r+1][c+1] = 0
			}
		}
	}
	
	if matrix[0][0] == 0 {
		for c := range len(matrix[0]) {
			matrix[0][c] = 0
		}
	}
	
	if setColZero {
		for r := range len(matrix) {
			matrix[r][0] = 0
		}
	}
}
