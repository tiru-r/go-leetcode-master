package pascals_triangle_118

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	result := make([][]int, numRows)
	
	for i := range numRows {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		
		result[i] = row
	}

	return result
}
