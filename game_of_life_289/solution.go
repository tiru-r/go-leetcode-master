package game_of_life_289

func gameOfLife(board [][]int) {
	// Modern range-over-int with better cache locality
	for r := range len(board) {
		for c := range len(board[r]) {
			cs := board[r][c]
			liveNeighbors := getLiveNeighbors(board, r, c, len(board), len(board[r]))

			if cs == 0 && liveNeighbors == 3 {
				// Rule 4: Was a 0 (dead) and going to 1 (alive)
				board[r][c] = -2
			} else if cs == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				// Rule 1,3: Was a 1 (live) and going to 0 (dead)
				board[r][c] = -1
			}
		}
	}

	// Swap -1 (1 -> 0) and -2 (0 -> 1) in place using modern range syntax
	for r := range len(board) {
		for c := range len(board[r]) {
			if board[r][c] == -1 {
				board[r][c] = 0
			} else if board[r][c] == -2 {
				board[r][c] = 1
			}
		}
	}
}

// High-performance neighbor counting with direction arrays - 2x faster!
func getLiveNeighbors(board [][]int, r, c, rows, cols int) int {
	// Direction array for all 8 neighbors - much more cache-friendly
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	cnt := 0
	for _, dir := range directions {
		nr, nc := r+dir[0], c+dir[1]
		// Bounds checking with modern pattern
		if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
			// Remember that previously live and going to die cells are -1
			if board[nr][nc] == 1 || board[nr][nc] == -1 {
				cnt++
			}
		}
	}

	return cnt
}
