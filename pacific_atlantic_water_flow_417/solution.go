package pacific_atlantic_water_flow_417

// pacificAtlantic returns all coordinates [r, c] that can flow to both oceans.
// Water flows from a cell to an adjacent cell if the adjacent cell's height
// is less than or equal to the current cell's height.
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])

	// reachable flags for Pacific and Atlantic
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// DFS from the four borders
	var dfs func(r, c int, prevHeight int, ocean [][]bool)
	dfs = func(r, c int, prevHeight int, ocean [][]bool) {
		if r < 0 || r >= m || c < 0 || c >= n ||
			ocean[r][c] || heights[r][c] < prevHeight {
			return
		}
		ocean[r][c] = true
		// 4-directional neighbors
		for _, d := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			dfs(r+d[0], c+d[1], heights[r][c], ocean)
		}
	}

	// Pacific borders: top row and left column
	for i := 0; i < m; i++ {
		dfs(i, 0, heights[i][0], pacific)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, heights[0][j], pacific)
	}

	// Atlantic borders: bottom row and right column
	for i := 0; i < m; i++ {
		dfs(i, n-1, heights[i][n-1], atlantic)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, heights[m-1][j], atlantic)
	}

	// Collect results
	var result [][]int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}
