package design_tic_tac_toe_348

import "testing"

func TestTicTacToe(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		moves    [][3]int // {row, col, player}
		winIndex int      // index of move that should return winner (or -1 for none)
		winner   int
	}{
		{
			name:     "3x3 player 1 wins on row",
			size:     3,
			moves:    [][3]int{{0, 0, 1}, {1, 0, 2}, {0, 1, 1}, {1, 1, 2}, {0, 2, 1}},
			winIndex: 4,
			winner:   1,
		},
		{
			name:     "3x3 player 2 wins on anti-diagonal",
			size:     3,
			moves:    [][3]int{{0, 0, 1}, {0, 2, 2}, {1, 0, 1}, {1, 1, 2}, {2, 1, 1}, {2, 0, 2}},
			winIndex: 5,
			winner:   2,
		},
		{
			name:     "3x3 player 1 wins on main diagonal",
			size:     3,
			moves:    [][3]int{{0, 0, 1}, {0, 1, 2}, {1, 1, 1}, {0, 2, 2}, {2, 2, 1}},
			winIndex: 4,
			winner:   1,
		},
		{
			name:     "3x3 player 2 wins on column",
			size:     3,
			moves:    [][3]int{{0, 0, 1}, {0, 1, 2}, {1, 0, 1}, {1, 1, 2}, {0, 2, 1}, {2, 1, 2}},
			winIndex: 5,
			winner:   2,
		},
		{
			name:     "4x4 player 1 wins on row",
			size:     4,
			moves:    [][3]int{{0, 0, 1}, {1, 0, 2}, {0, 1, 1}, {1, 1, 2}, {0, 2, 1}, {1, 2, 2}, {0, 3, 1}},
			winIndex: 6,
			winner:   1,
		},
		{
			name:     "3x3 draw",
			size:     3,
			moves:    [][3]int{{0, 0, 1}, {0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 1, 1}, {1, 2, 2}, {2, 0, 2}, {2, 1, 1}, {2, 2, 2}},
			winIndex: -1,
			winner:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game := NewTicTacToe(tt.size)
			for i, move := range tt.moves {
				got := game.Move(move[0], move[1], move[2])
				if i == tt.winIndex && got != tt.winner {
					t.Fatalf("move %d: expected winner %d, got %d", i, tt.winner, got)
				}
				if i != tt.winIndex && got != 0 {
					t.Fatalf("move %d: unexpected winner %d", i, got)
				}
			}
		})
	}
}
