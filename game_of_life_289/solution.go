package game_of_life_289

const (
	liveBit = 1 // current state bit
	nextBit = 2 // next state bit
	numDirections = 8
)

// gameOfLife updates board in-place according to Conway's rules.
// Encoding:  bit-0 is “current”, bit-1 is “next”.
//
//	0 0  dead → dead
//	0 1  dead → alive
//	1 0  alive → dead
//	1 1  alive → alive
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	
	// Direction vectors for 8 neighbors
	directions := [numDirections][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	// Pass 1: mark next state in bit-1
	for r := range m {
		for c := range n {
			live := 0
			for _, dir := range directions {
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && board[nr][nc]&liveBit != 0 {
					live++
				}
			}
			alive := board[r][c]&liveBit != 0
			if (alive && (live == 2 || live == 3)) || (!alive && live == 3) {
				board[r][c] |= nextBit
			}
		}
	}

	// Pass 2: shift bit-1 to bit-0
	for r := range m {
		for c := range n {
			board[r][c] >>= 1
		}
	}
}
