package spiral_matrix_54

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	
	const maxDimension = 10000 // Support matrices up to 10000x10000
	seen := make(map[int]struct{})
	var spiral []int
	seenStep := false
	dir := "east"
	var r, c int
	
	for !seenStep && !isOutOfBounds(matrix, r, c) {
		seen[r*maxDimension+c] = struct{}{}
		spiral = append(spiral, matrix[r][c])

		switch dir {
		case "north":
			r--
		case "east":
			c++
		case "south":
			r++
		case "west":
			c--
		}

		_, seenStep = seen[r*maxDimension+c]
		if isOutOfBounds(matrix, r, c) || seenStep {
			switch dir {
			case "north":
				r++
				dir = "east"
				c++
			case "east":
				c--
				dir = "south"
				r++
			case "south":
				r--
				dir = "west"
				c--
			case "west":
				c++
				dir = "north"
				r--
			}
		}

		_, seenStep = seen[r*maxDimension+c]
	}

	return spiral
}

func isOutOfBounds(m [][]int, r, c int) bool {
	return r < 0 || r >= len(m) ||
		c < 0 || (r >= 0 && r < len(m) && c >= len(m[r]))
}
