package coin_change_322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1: Basic change",
			args: args{
				coins:  []int{1, 2},
				amount: 3,
			},
			want: 2, // 1 + 2 or 1 + 1 + 1 (min is 2)
		},
		{
			name: "Example 2: Larger amount with multiple denominations",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3, // 5 + 5 + 1
		},
		{
			name: "Example 3: No solution possible",
			args: args{
				coins:  []int{3},
				amount: 1,
			},
			want: -1,
		},
		{
			name: "Example 4: Amount is zero",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 0,
			},
			want: 0,
		},
		{
			name: "Example 5: Complex case with various coins",
			args: args{
				coins:  []int{186, 419, 83, 408},
				amount: 6249,
			},
			want: 20,
		},
		{
			name: "Example 6: Single coin matches amount",
			args: args{
				coins:  []int{5},
				amount: 5,
			},
			want: 1,
		},
		{
			name: "Example 7: Coins larger than amount",
			args: args{
				coins:  []int{10, 20},
				amount: 5,
			},
			want: -1,
		},
		{
			name: "Example 8: Duplicate coins (should still work)",
			args: args{
				coins:  []int{1, 1, 2},
				amount: 3,
			},
			want: 2, // 1 + 2
		},
		{
			name: "Example 9: Large amount, single coin",
			args: args{
				coins:  []int{2},
				amount: 100,
			},
			want: 50,
		},
		{
			name: "Example 10: No coins provided",
			args: args{
				coins:  []int{},
				amount: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, coinChange(tt.args.coins, tt.args.amount))
		})
	}
}
