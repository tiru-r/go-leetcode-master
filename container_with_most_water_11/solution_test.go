package container_with_most_water_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "standard_case",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "two_lines", 
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "four_lines",
			height: []int{1, 2, 4, 3},
			want:   4,
		},
		{
			name:   "empty_array",
			height: []int{},
			want:   0,
		},
		{
			name:   "single_element",
			height: []int{5},
			want:   0,
		},
		{
			name:   "all_same_height",
			height: []int{3, 3, 3, 3, 3},
			want:   12,
		},
		{
			name:   "increasing_heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.height))
		})
	}
}
