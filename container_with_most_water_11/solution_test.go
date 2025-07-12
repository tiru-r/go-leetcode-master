package container_with_most_water_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_maxArea tests the maxArea function with various input scenarios to ensure
// it correctly calculates the maximum area of water that can be trapped between lines.
func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard_case",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49, // Max area from height[1]=8 and height[8]=7 with width=7.
		},
		{
			name: "two_lines",
			args: args{
				height: []int{1, 1},
			},
			want: 1, // Only two lines, area = min(1,1) * 1 = 1.
		},
		{
			name: "four_lines",
			args: args{
				height: []int{1, 2, 4, 3},
			},
			want: 4, // Max area from height[1]=2 and height[2]=4 with width=2.
		},
		{
			name: "empty_array",
			args: args{
				height: []int{},
			},
			want: 0, // No lines, no area possible.
		},
		{
			name: "single_element",
			args: args{
				height: []int{5},
			},
			want: 0, // Single line can't form a container.
		},
		{
			name: "all_same_height",
			args: args{
				height: []int{3, 3, 3, 3, 3},
			},
			want: 12, // Max area from height[0]=3 and height[4]=3 with width=4.
		},
		{
			name: "increasing_heights",
			args: args{
				height: []int{1, 2, 3, 4, 5},
			},
			want: 6, // Max area from height[0]=1 and height[4]=5 with width=4.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.args.height), "Failed for input: %v", tt.args.height)
		})
	}
}
