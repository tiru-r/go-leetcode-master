package pacific_atlantic_water_flow_417

// Direction vectors for 4-directional movement (up, down, left, right)
var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// pacificAtlantic finds cells from which water can flow to both Pacific and Atlantic oceans
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])
	// Track cells reachable from each ocean using reverse DFS
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range m {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// Start DFS from all border cells of both oceans
	// Pacific: top and left borders
	for i := range m {
		dfs(heights, pacific, i, 0, m, n) // Left border
	}
	for j := range n {
		dfs(heights, pacific, 0, j, m, n) // Top border
	}

	// Atlantic: bottom and right borders
	for i := range m {
		dfs(heights, atlantic, i, n-1, m, n) // Right border
	}
	for j := range n {
		dfs(heights, atlantic, m-1, j, m, n) // Bottom border
	}

	// Collect cells reachable from both oceans
	// Pre-allocate with estimated capacity (worst case: all cells)
	result := make([][]int, 0, min(m*n/4, 100)) // Conservative estimate
	for row := range m {
		for col := range n {
			if pacific[row][col] && atlantic[row][col] {
				result = append(result, []int{row, col})
			}
		}
	}
	return result
}

// dfs performs depth-first search to mark cells reachable from ocean borders
// Water flows from higher or equal elevation to current cell
func dfs(heights [][]int, ocean [][]bool, row, col, m, n int) {
	ocean[row][col] = true
	for _, dir := range directions {
		nextRow, nextCol := row+dir[0], col+dir[1]
		if nextRow >= 0 && nextRow < m && nextCol >= 0 && nextCol < n &&
			!ocean[nextRow][nextCol] && heights[nextRow][nextCol] >= heights[row][col] {
			dfs(heights, ocean, nextRow, nextCol, m, n)
		}
	}
}
