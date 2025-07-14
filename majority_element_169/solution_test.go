package majority_element_169

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 1},
		{"two same elements", []int{3, 3}, 3},
		{"simple majority", []int{3, 2, 3}, 3},
		{"complex case", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"long sequence", []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7}, 7},
		{"negative numbers", []int{-1, -1, -2, -1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, majorityElement(tt.nums))
		})
	}
}

func BenchmarkMajorityElement(b *testing.B) {
	nums := []int{3, 2, 3, 1, 3, 4, 3, 5, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}

func BenchmarkMajorityElementLarge(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		if i%2 == 0 {
			nums[i] = 1 // majority element
		} else {
			nums[i] = rand.Intn(100) + 2
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		majorityElement(nums)
	}
}
