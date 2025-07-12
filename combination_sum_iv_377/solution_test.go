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
		{"large target", []int{1, 2}, 6, 13}, // corrected canonical value
		{"all too big", []int{10, 20}, 5, 0},
		{"negative target", []int{1, 2}, -3, 0},
		{"zero coins", []int{}, 0, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// 1. Drop-in wrapper
			assert.Equal(t, tc.want, CombinationSum(tc.nums, tc.target))

			// 2. In-place version with fresh buffer
			buf := make([]int, tc.target+1)
			assert.Equal(t, tc.want, CombinationSumInPlace(tc.nums, tc.target, buf))
		})
	}
}

// Benchmark the wrapper (uses sync.Pool)
func BenchmarkCombinationSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CombinationSum([]int{1, 2, 3}, 32)
	}
}

// Benchmark the in-place variant
func BenchmarkCombinationSum4InPlace(b *testing.B) {
	buf := make([]int, 33)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSumInPlace([]int{1, 2, 3}, 32, buf)
	}
}
