package game_of_life_289

import (
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		// 1. trivial sizes
		{"empty", [][]int{}, [][]int{}},
		{"1x1 dead", [][]int{{0}}, [][]int{{0}}},
		{"1x1 alive", [][]int{{1}}, [][]int{{0}}},
		{"1x2 all dead", [][]int{{0, 0}}, [][]int{{0, 0}}},
		{"2x1 all dead", [][]int{{0}, {0}}, [][]int{{0}, {0}}},
		{"2x2 all dead", [][]int{{0, 0}, {0, 0}}, [][]int{{0, 0}, {0, 0}}},
		{"2x2 all alive", [][]int{{1, 1}, {1, 1}}, [][]int{{1, 1}, {1, 1}}},
		{"2x2 blinker", [][]int{{1, 0}, {1, 0}}, [][]int{{0, 0}, {0, 0}}},

		// 2. canonical 3×3 examples
		{"block stays", [][]int{{1, 1}, {1, 1}}, [][]int{{1, 1}, {1, 1}}},
		{"blinker vertical -> horizontal", [][]int{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}}, [][]int{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}}},
		{"LeetCode sample", [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}},

		// 3. edge rules
		{"live with 1 neighbor dies", [][]int{{1, 0}, {0, 1}}, [][]int{{0, 0}, {0, 0}}},
		{"live with 4 neighbors dies", [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{"dead with 3 neighbors becomes alive", [][]int{{0, 1, 0}, {1, 1, 0}, {0, 0, 0}}, [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 0}}},

		// 4. larger boards
		{"4x4 simple", [][]int{{0, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 0}, {1, 0, 0, 1}}, [][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}},
		{"5x5 oscillator", [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
		}, [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := deepCopy(tt.board)
			gameOfLife(board)
			if !reflect.DeepEqual(board, tt.want) {
				t.Fatalf("%s:\ninput %v\nwant  %v\ngot   %v",
					tt.name, tt.board, tt.want, board)
			}
		})
	}
}

/* ---------- helper: deep copy board ---------- */
func deepCopy(b [][]int) [][]int {
	out := make([][]int, len(b))
	for i := range b {
		out[i] = append([]int(nil), b[i]...)
	}
	return out
}
