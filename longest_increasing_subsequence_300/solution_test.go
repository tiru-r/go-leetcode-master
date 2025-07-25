package longest_increasing_subsequence_300

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"strictly increasing", []int{1, 2, 3, 4, 5}, 5},
		{"strictly decreasing", []int{5, 4, 3, 2, 1}, 1},
		{"duplicates allowed", []int{2, 2, 2}, 1},
		{"example 1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"example 2", []int{1, 3, 2}, 2},
		{"negative numbers", []int{-2, -1, -3, -1, -2}, 2},
		{"mixed", []int{3, 5, 6, 2, 5, 4, 19, 5, 6, 7, 12}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLIS(tt.nums))
		})
	}
}

func BenchmarkLengthOfLIS(b *testing.B) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18, 4, 6, 11, 15, 20}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}

func BenchmarkLengthOfLISLarge(b *testing.B) {
	nums := make([]int, 2500)
	for i := range nums {
		nums[i] = rand.Intn(10000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLIS(nums)
	}
}
