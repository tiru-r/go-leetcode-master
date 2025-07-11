package spiral_matrix_54

// Ultra-efficient O(1) space, O(m*n) time boundary-based spiral traversal
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n) // Pre-allocate exact capacity

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

// Old inefficient O(m*n) space solution - DO NOT USE (kept for comparison)
func spiralOrderInefficient(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	const maxDimension = 10000 // Wastes memory with large maps
	seen := make(map[int]struct{})
	var spiral []int
	seenStep := false
	var r, c int
	
	for !seenStep && !isOutOfBoundsOld(matrix, r, c) {
		seen[r*maxDimension+c] = struct{}{}
		spiral = append(spiral, matrix[r][c])
		// Inefficient direction-based traversal omitted for brevity
		break // Prevent infinite loop in this stub
	}
	return spiral
}

func isOutOfBoundsOld(m [][]int, r, c int) bool {
	return r < 0 || r >= len(m) ||
		c < 0 || (r >= 0 && r < len(m) && c >= len(m[r]))
}
