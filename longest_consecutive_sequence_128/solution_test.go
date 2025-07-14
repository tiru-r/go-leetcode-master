package longest_consecutive_sequence_128

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"classic example", []int{100, 4, 200, 1, 3, 2}, 4},
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 1},
		{"duplicates", []int{1, 2, 2, 2, 3}, 3},
		{"negative numbers", []int{-2, -1, 0, 1, 2}, 5},
		{"no sequence", []int{9, 1, 4, 7, 3}, 2},
		{"reverse order", []int{5, 4, 3, 2, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestConsecutive(tt.nums))
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	nums := []int{100, 4, 200, 1, 3, 2, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestConsecutive(nums)
	}
}

func BenchmarkLongestConsecutiveLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = rand.Intn(20000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestConsecutive(nums)
	}
}
