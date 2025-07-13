package game_of_life_289

// gameOfLife applies Conway's Game of Life rules in-place
// Optimized: O(m*n) time, O(1) space using state encoding
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	
	rows, cols := len(board), len(board[0])
	
	// Pass 1: Mark transitions using state encoding
	// -1: was alive, dies  -2: was dead, becomes alive
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			neighbors := countLiveNeighbors(board, r, c, rows, cols)
			
			// Apply Game of Life rules
			if board[r][c] == 1 {
				// Live cell dies if < 2 or > 3 neighbors
				if neighbors < 2 || neighbors > 3 {
					board[r][c] = -1
				}
			} else {
				// Dead cell becomes alive with exactly 3 neighbors
				if neighbors == 3 {
					board[r][c] = -2
				}
			}
		}
	}
	
	// Pass 2: Apply transitions using bitwise operations
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Efficient state restoration
			if board[r][c] == -1 {
				board[r][c] = 0
			} else if board[r][c] == -2 {
				board[r][c] = 1
			}
		}
	}
}

// countLiveNeighbors efficiently counts live neighbors using direction vectors
func countLiveNeighbors(board [][]int, r, c, rows, cols int) int {
	// Pre-computed neighbor offsets for cache efficiency
	neighbors := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}
	
	count := 0
	for _, offset := range neighbors {
		nr, nc := r+offset[0], c+offset[1]
		
		// Bounds check with early termination
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			// Count original live cells (1 or -1 for dying)
			if board[nr][nc] == 1 || board[nr][nc] == -1 {
				count++
			}
		}
	}
	
	return count
}
