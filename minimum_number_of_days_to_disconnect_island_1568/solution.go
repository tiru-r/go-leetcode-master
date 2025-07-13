package minimum_number_of_days_to_disconnect_island_1568

// MinDays finds minimum days to disconnect the island
// Time: O(m*n*(m*n)), Space: O(m*n)
func MinDays(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	
	// Check if already disconnected
	if countIslands(grid) != 1 {
		return 0
	}
	
	// Try removing each land cell (1 day)
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if countIslands(grid) != 1 {
					return 1
				}
				grid[i][j] = 1 // backtrack
			}
		}
	}
	
	// Maximum days needed is 2 (proven by theory)
	return 2
}

// countIslands counts number of connected island components
func countIslands(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	
	count := 0
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 && !visited[i][j] {
				dfs(grid, visited, i, j)
				count++
			}
		}
	}
	
	return count
}

// dfs performs depth-first search to mark connected land cells
func dfs(grid [][]int, visited [][]bool, i, j int) {
	m, n := len(grid), len(grid[0])
	
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == 0 {
		return
	}
	
	visited[i][j] = true
	
	// Visit all 4 directions
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		dfs(grid, visited, i+dir[0], j+dir[1])
	}
}