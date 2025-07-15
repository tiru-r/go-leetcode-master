package pacific_atlantic_water_flow_417

var directions = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	m, n := len(heights), len(heights[0])
	pacific := makeGrid(m, n)
	atlantic := makeGrid(m, n)

	for i := range m {
		dfs(heights, pacific, i, 0, m, n)
		dfs(heights, atlantic, i, n-1, m, n)
	}

	for j := range n {
		dfs(heights, pacific, 0, j, m, n)
		dfs(heights, atlantic, m-1, j, m, n)
	}

	result := make([][]int, 0)
	for r := range m {
		for c := range n {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

func makeGrid(m, n int) [][]bool {
	grid := make([][]bool, m)
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	return grid
}

func dfs(heights [][]int, ocean [][]bool, r, c, m, n int) {
	ocean[r][c] = true

	for _, dir := range directions {
		nr, nc := r+dir[0], c+dir[1]
		if nr >= 0 && nr < m && nc >= 0 && nc < n &&
			!ocean[nr][nc] && heights[nr][nc] >= heights[r][c] {
			dfs(heights, ocean, nr, nc, m, n)
		}
	}
}
