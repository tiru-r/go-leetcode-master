package combination_sum_iv_377

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum4(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"standard", []int{1, 2, 3}, 4, 7},
		{"impossible", []int{3}, 5, 0},
		{"empty nums", []int{}, 1, 0},
		{"target zero", []int{1, 2, 3}, 0, 1},
		{"single exact", []int{5}, 5, 1},
		{"large target", []int{1, 2}, 6, 13},
		{"all too big", []int{10, 20}, 5, 0},
		{"negative target", []int{1, 2}, -3, 0},
		{"zero coins", []int{}, 0, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, combinationSum4(tc.nums, tc.target))
		})
	}
}

func BenchmarkCombinationSum4(b *testing.B) {
	for range b.N {
		combinationSum4([]int{1, 2, 3}, 32)
	}
}
