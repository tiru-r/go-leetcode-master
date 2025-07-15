package number_of_islands_200

const (
	water = 0
	land  = 1
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	m, n := len(grid), len(grid[0])

	for r := range m {
		for c := range n {
			if grid[r][c] == land {
				count++
				dfs(grid, r, c, m, n)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c, m, n int) {
	if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != land {
		return
	}

	grid[r][c] = water

	dfs(grid, r-1, c, m, n)
	dfs(grid, r+1, c, m, n)
	dfs(grid, r, c-1, m, n)
	dfs(grid, r, c+1, m, n)
}
