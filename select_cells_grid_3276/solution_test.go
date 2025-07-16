package select_cells_grid_3276

import (
	"testing"
)

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "empty grid",
			grid: [][]int{},
			want: 0,
		},
		{
			name: "empty row",
			grid: [][]int{{}},
			want: 0,
		},
		{
			name: "single cell",
			grid: [][]int{{5}},
			want: 5,
		},
		{
			name: "all values equal",
			grid: [][]int{{7, 7}, {7, 7}},
			want: 7,
		},
		{
			name: "no duplicates in grid",
			grid: [][]int{{3, 1}, {4, 2}},
			want: 4 + 3, // rows 0 & 1, values 4 and 3
		},
		{
			name: "duplicates across rows",
			grid: [][]int{{5, 9}, {5, 7}},
			want: 9 + 7, // rows 0 & 1, values 9 and 7
		},
		{
			name: "must skip duplicate value in same row",
			grid: [][]int{{8, 8}, {3, 4}},
			want: 8 + 4,
		},
		{
			name: "largest values force overlap on rows",
			grid: [][]int{{10, 3}, {10, 2}},
			want: 12, // 10 from either row + 2 from second row
		},
		{
			name: "three rows with overlaps",
			grid: [][]int{{6, 5}, {5, 4}, {6, 3}},
			want: 6 + 5 + 3, // rows 0,1,2 values 6,5,3
		},
		{
			name: "negative values allowed",
			grid: [][]int{{-1, -2}, {-3, -4}},
			want: -1, // largest single value
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.grid); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
