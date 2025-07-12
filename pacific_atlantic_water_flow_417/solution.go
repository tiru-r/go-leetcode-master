package pacific_atlantic_water_flow_417

// pacificAtlantic returns all coordinates [r,c] from which water can reach
// both the Pacific and Atlantic oceans.
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// DFS helper
	var dfs func(visited [][]bool, r, c, prev int)
	dfs = func(visited [][]bool, r, c, prev int) {
		if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || heights[r][c] < prev {
			return
		}
		visited[r][c] = true
		dfs(visited, r-1, c, heights[r][c])
		dfs(visited, r+1, c, heights[r][c])
		dfs(visited, r, c-1, heights[r][c])
		dfs(visited, r, c+1, heights[r][c])
	}

	// Start DFS from Pacific borders (top & left)
	for r := 0; r < m; r++ {
		dfs(pacific, r, 0, heights[r][0])
	}
	for c := 0; c < n; c++ {
		dfs(pacific, 0, c, heights[0][c])
	}

	// Start DFS from Atlantic borders (bottom & right)
	for r := 0; r < m; r++ {
		dfs(atlantic, r, n-1, heights[r][n-1])
	}
	for c := 0; c < n; c++ {
		dfs(atlantic, m-1, c, heights[m-1][c])
	}

	// Collect cells reachable from both oceans
	var res [][]int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}
	return res
}
