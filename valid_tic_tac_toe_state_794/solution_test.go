package valid_tic_tac_toe_state_794

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validTicTacToe(t *testing.T) {
	tests := []struct {
		name  string
		board []string
		want  bool
	}{
		{"too many O's", []string{"O  ", "   ", "   "}, false},
		{"X out of turn", []string{"XOX", " X ", "   "}, false},
		{"both win", []string{"XXX", "   ", "OOO"}, false},
		{"X wins on correct turn", []string{"XXX", " O ", "  O"}, true},
		{"O wins on correct turn", []string{"OOO", "XX ", "  X"}, true},
		{"draw", []string{"XOX", "O O", "XOX"}, true},
		{"empty board", []string{"   ", "   ", "   "}, true},
		{"full board draw", []string{"XOX", "XXO", "OXO"}, true},
		{"full board draw alt", []string{"XOX", "XXO", "OOX"}, true},
		{"X wins diagonal", []string{"XO ", "OX ", "  X"}, true},
		{"X wins row", []string{"XXX", "OOX", "OOX"}, true},
		{"X wins but wrong count", []string{"XXX", "XOO", "OO "}, false},
		{"single X", []string{"X  ", "   ", "   "}, true},
		{"single O", []string{"O  ", "   ", "   "}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validTicTacToe(tt.board))
		})
	}
}
