package spiral_matrix_54

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		for col := range right - left + 1 {
			result = append(result, matrix[top][left+col])
		}
		top++

		for row := range bottom - top + 1 {
			result = append(result, matrix[top+row][right])
		}
		right--

		if top <= bottom {
			for col := range right - left + 1 {
				result = append(result, matrix[bottom][right-col])
			}
			bottom--
		}

		if left <= right {
			for row := range bottom - top + 1 {
				result = append(result, matrix[bottom-row][left])
			}
			left++
		}
	}

	return result
}
