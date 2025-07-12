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
		{"empty", [][]int{}, nil},
		{"1x1", [][]int{{1}}, [][]int{{0, 0}}},
		{"2x2", [][]int{{1, 2}, {2, 1}}, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
		{"leetcode example",
			[][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pacificAtlantic(tt.grid)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
