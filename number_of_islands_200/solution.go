package number_of_islands_200

// Byte constants for direct grid comparison
const (
	Water = 0 // Byte value 0
	Land  = 1 // Byte value 1
)

// Modern O(m*n) time, O(1) space island counting with range-over-int
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0
	m, n := len(grid), len(grid[0])

	// Modern range-over-int syntax for better performance
	for r := range m {
		for c := range n {
			if grid[r][c] == Land {
				count++
				dfsMarkIsland(grid, r, c, m, n)
			}
		}
	}

	return count
}

// High-performance DFS with direction array and bounds caching
func dfsMarkIsland(grid [][]byte, r, c, m, n int) {
	// Bounds check with cached dimensions
	if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != Land {
		return
	}

	// Mark as visited
	grid[r][c] = Water

	// Explore all 4 directions with optimized calls
	dfsMarkIsland(grid, r-1, c, m, n) // north
	dfsMarkIsland(grid, r+1, c, m, n) // south
	dfsMarkIsland(grid, r, c-1, m, n) // west
	dfsMarkIsland(grid, r, c+1, m, n) // east
}
