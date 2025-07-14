package minimum_number_of_days_to_disconnect_island_1568

func minDays(grid [][]int) int {
	if countIslands(grid) != 1 {
		return 0
	}
	
	m, n := len(grid), len(grid[0])
	for i := range m {
		for j := range n {
			if grid[i][j] == 1 {
				grid[i][j] = 0
				if countIslands(grid) != 1 {
					return 1
				}
				grid[i][j] = 1
			}
		}
	}
	
	return 2
}

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

func dfs(grid [][]int, visited [][]bool, i, j int) {
	m, n := len(grid), len(grid[0])
	
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || grid[i][j] == 0 {
		return
	}
	
	visited[i][j] = true
	
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		dfs(grid, visited, i+dir[0], j+dir[1])
	}
}