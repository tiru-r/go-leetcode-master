package spiral_matrix_54

// Ultra-efficient O(1) space, O(m*n) time boundary-based spiral traversal
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	totalElements := m * n
	result := make([]int, 0, totalElements) // Pre-allocate exact capacity

	// Boundary variables - much more efficient than tracking visited cells
	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// Traverse right across top row
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// Traverse down right column
		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// Traverse left across bottom row (if still valid)
		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		// Traverse up left column (if still valid)
		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}
