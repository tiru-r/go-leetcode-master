package valid_tic_tac_toe_state_794

const (
	boardSize = 3
	playerX   = 'X'
	playerO   = 'O'
	emptyCell = ' '
)

// validTicTacToe checks if a tic-tac-toe board represents a valid game state.
// Rules: X goes first, players alternate, at most one winner, correct move counts.
func validTicTacToe(board []string) bool {
	// Count pieces for each player
	xCount, oCount := 0, 0
	for _, row := range board {
		for _, cell := range row {
			switch cell {
			case playerX:
				xCount++
			case playerO:
				oCount++
			}
		}
	}

	// Validate move counts: X goes first, so xCount == oCount or xCount == oCount + 1
	if xCount != oCount && xCount != oCount+1 {
		return false
	}

	// Define all possible winning lines (rows, columns, diagonals)
	winningLines := [8][3][2]int{
		// Rows
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		// Columns
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		// Diagonals
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	// Check for winning lines
	xHasWin, oHasWin := false, false
	for _, line := range winningLines {
		first := board[line[0][0]][line[0][1]]
		second := board[line[1][0]][line[1][1]]
		third := board[line[2][0]][line[2][1]]
		
		// Check if all three cells have the same non-empty value
		if first != emptyCell && first == second && second == third {
			switch first {
			case playerX:
				xHasWin = true
			case playerO:
				oHasWin = true
			}
			
			// Early termination: if both have winning lines, invalid state
			if xHasWin && oHasWin {
				return false
			}
		}
	}

	// Validate game state based on winners
	if xHasWin && oHasWin {
		return false // Both players cannot have winning lines
	}
	if xHasWin && xCount != oCount+1 {
		return false // If X wins, X must have made the last move
	}
	if oHasWin && xCount != oCount {
		return false // If O wins, O must have made the last move
	}
	
	return true // Valid game state
}
