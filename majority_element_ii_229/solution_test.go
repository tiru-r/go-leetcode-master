package majority_element_ii_229

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"empty slice", []int{}, []int{}},
		{"single element", []int{1}, []int{1}},
		{"two elements same", []int{1, 1}, []int{1}},
		{"two elements different", []int{1, 2}, []int{1, 2}}, // Both appear 1 time > 2/3=0
		{"simple majority", []int{3, 2, 3}, []int{3}},
		{"two majorities", []int{1, 1, 1, 3, 3, 2, 2, 2}, []int{1, 2}},
		{"no majority", []int{1, 2, 3, 4, 5}, []int{}},
		{"three elements", []int{1, 2, 3}, []int{}}, // Each appears 1 time, not > 3/3=1
		{"negative numbers", []int{-1, -1, -1, -2, -2, -3}, []int{-1}},
		{"duplicates", []int{0, 0, 0}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.nums)
			slices.Sort(result)
			slices.Sort(tt.want)
			assert.Equal(t, tt.want, result)
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := []int{1, 1, 1, 3, 3, 2, 2, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}

func BenchmarkMajorityElementLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		switch i % 3 {
		case 0:
			nums[i] = 1 // majority element 1
		case 1:
			nums[i] = 2 // majority element 2
		default:
			nums[i] = rand.Intn(100) + 10 // random non-majority
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}

func BenchmarkMajorityElementNoMajority(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i // all unique, no majority
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}
