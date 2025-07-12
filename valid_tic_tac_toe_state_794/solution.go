package valid_tic_tac_toe_state_794

// validTicTacToe returns true iff the board is legal.
// board[i] is exactly three runes long and contains only {'X','O',' '}.
func validTicTacToe(board []string) bool {
	x, o := 0, 0
	for _, row := range board {
		for _, c := range row {
			switch c {
			case 'X':
				x++
			case 'O':
				o++
			}
		}
	}

	// Turn-order rule
	if x != o && x != o+1 {
		return false
	}

	// Pre-computed magic indices for the 8 lines.
	lines := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // cols
		{0, 4, 8}, {2, 4, 6}, // diagonals
	}

	flat := board[0] + board[1] + board[2]
	xWin, oWin := false, false
	for _, ln := range lines {
		c0, c1, c2 := flat[ln[0]], flat[ln[1]], flat[ln[2]]
		if c0 == c1 && c1 == c2 {
			switch c0 {
			case 'X':
				xWin = true
			case 'O':
				oWin = true
			}
		}
	}

	// Mutual win impossible
	if xWin && oWin {
		return false
	}
	// If X wins, X must have played last
	if xWin && x != o+1 {
		return false
	}
	// If O wins, O must have played last
	if oWin && x != o {
		return false
	}
	return true
}
