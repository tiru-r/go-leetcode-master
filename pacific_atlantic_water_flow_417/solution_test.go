package pacific_atlantic_water_flow_417

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{"empty grid", [][]int{}, nil},
		{"empty row", [][]int{{}}, nil},
		{"single cell", [][]int{{5}}, [][]int{{0, 0}}},
		{"2x2 all same", [][]int{{1, 1}, {1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
		{"2x2 slope", [][]int{{1, 2}, {2, 1}}, [][]int{{0, 1}, {1, 0}}},
		{"leetcode example",
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		{"3x3 monotone",
			[][]int{{3, 3, 3}, {3, 3, 3}, {3, 3, 3}},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{"all lower than neighbours",
			[][]int{{5, 4, 3}, {4, 3, 2}, {3, 2, 1}},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pacificAtlantic(tt.grid)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
