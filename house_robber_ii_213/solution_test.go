package house_robber_ii_213

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// 1. LeetCode official samples
		{"LC 2-3-2", []int{2, 3, 2}, 3},
		{"LC 1-2-3-1", []int{1, 2, 3, 1}, 4},
		{"LC 1-2-3", []int{1, 2, 3}, 3},
		{"LC 1-2-1-1", []int{1, 2, 1, 1}, 3},

		// 2. Boundary sizes
		{"empty", []int{}, 0},
		{"single house", []int{5}, 5},
		{"two houses", []int{2, 1}, 2},
		{"two equal", []int{3, 3}, 3},

		// 3. All zeroes
		{"all zeroes", []int{0, 0, 0, 0}, 0},

		// 4. Monotonic patterns
		{"increasing", []int{1, 2, 3, 4, 5}, 8}, // 2+4 (skip 1 and 5)
		{"decreasing", []int{5, 4, 3, 2, 1}, 8}, // 5+3 (skip last, can't rob 5+3+1 due to circular)

		// 5. Alternating high/low
		{"alternating", []int{10, 1, 10, 1, 10}, 20},

		// 6. Large input (performance smoke)
		{"large 100", make([]int, 100), 0}, // all zeroes
		{"large 100 random", func() []int {
			out := make([]int, 100)
			for i := range out {
				out[i] = i + 1 // 1..100
			}
			return out
		}(), 2550}, // solution = 50+52+...+100 = 2550

		// 7. Negative numbers (constraint is non-negative, but test anyway)
		{"negative not allowed", []int{-1, -2, -3}, 0}, // constraint violation, but test
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rob(tt.nums))
		})
	}
}
