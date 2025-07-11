package combination_sum_iv_377

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard case with multiple numbers",
			args: args{
				nums:   []int{1, 2, 3},
				target: 4,
			},
			want: 7, // Combinations: [1,1,1,1], [1,1,2], [1,2,1], [2,1,1], [2,2], [1,3], [3,1]
		},
		{
			name: "impossible sum with single number",
			args: args{
				nums:   []int{3},
				target: 5,
			},
			want: 0, // No way to make 5 with only 3
		},
		{
			name: "empty nums array",
			args: args{
				nums:   []int{},
				target: 1,
			},
			want: 0, // Empty array cannot form any positive sum
		},
		{
			name: "target zero",
			args: args{
				nums:   []int{1, 2, 3},
				target: 0,
			},
			want: 1, // One way: empty combination
		},
		{
			name: "single number equals target",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: 1, // Only one combination: [5]
		},
		{
			name: "large target with multiple numbers",
			args: args{
				nums:   []int{1, 2},
				target: 6,
			},
			want: 8, // Combinations: [1,1,1,1,1,1], [1,1,1,1,2], [1,1,1,2,1], [1,1,2,1,1], [1,2,1,1,1], [2,1,1,1,1], [1,2,2,1], [2,1,2,1]
		},
		{
			name: "all numbers larger than target",
			args: args{
				nums:   []int{10, 20, 30},
				target: 5,
			},
			want: 0, // No combinations possible
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, combinationSum4(tt.args.nums, tt.args.target))
		})
	}
}
