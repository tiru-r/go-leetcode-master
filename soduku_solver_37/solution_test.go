package soduku_solver_37

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	// Helper function to create a board from string array
	makeBoard := func(strs []string) [][]byte {
		board := make([][]byte, 9)
		for i, s := range strs {
			board[i] = []byte(s)
		}
		return board
	}

	// Helper function to validate a solved Sudoku board
	isValidSudoku := func(board [][]byte) bool {
		var rowMask, colMask, boxMask [9]uint16
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if board[r][c] == '.' {
					return false
				}
				d := uint16(board[r][c] - '0')
				if d < 1 || d > 9 {
					return false
				}
				bit := uint16(1) << (d - 1)
				bi := (r/3)*3 + c/3
				if rowMask[r]&bit != 0 || colMask[c]&bit != 0 || boxMask[bi]&bit != 0 {
					return false
				}
				rowMask[r] |= bit
				colMask[c] |= bit
				boxMask[bi] |= bit
			}
		}
		return true
	}

	tests := []struct {
		name        string
		input       [][]byte
		expectSolve bool
	}{
		{
			name: "Example board",
			input: makeBoard([]string{
				"53..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"....8..79",
			}),
			expectSolve: true,
		},
		{
			name: "Empty board",
			input: makeBoard([]string{
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			}),
			expectSolve: true,
		},
		{
			name: "Invalid board (unsolvable)",
			input: makeBoard([]string{
				"53..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"555.8..79", // Row 9 has three 5s, making it unsolvable
			}),
			expectSolve: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input board
			original := make([][]byte, 9)
			for i := range tt.input {
				original[i] = make([]byte, 9)
				copy(original[i], tt.input[i])
			}

			// Solve the board
			solveSudoku(tt.input)

			// If we expect a solution, verify it's valid
			if tt.expectSolve {
				if !isValidSudoku(tt.input) {
					t.Errorf("Test %s: Expected a valid Sudoku solution, but got an invalid board", tt.name)
				}
			} else {
				// For unsolvable board, verify it wasn't modified to an invalid state
				// or check that it remains unsolved (still has dots)
				hasDots := false
				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						if tt.input[i][j] == '.' {
							hasDots = true
							break
						}
					}
				}
				if isValidSudoku(tt.input) && !hasDots {
					t.Errorf("Test %s: Expected unsolvable board, but got a valid solution", tt.name)
				}
			}

			// Verify input board integrity (solver shouldn't modify input structure)
			if !reflect.DeepEqual(len(tt.input), 9) || !reflect.DeepEqual(len(tt.input[0]), 9) {
				t.Errorf("Test %s: Solver modified board structure", tt.name)
			}
		})
	}
}
