package trapping_rain_water_42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"empty", []int{}, 0},
		{"single bar", []int{5}, 0},
		{"two bars", []int{2, 1}, 0},
		{"classic example", []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{"flat plateau", []int{3, 3, 3, 3}, 0},
		{"plateau with dip", []int{4, 2, 3, 4, 3, 2, 4}, 9},
		{"ascending", []int{1, 2, 3, 4, 5}, 0},
		{"descending", []int{5, 4, 3, 2, 1}, 0},
		{"all zeros", []int{0, 0, 0, 0}, 0},
		{"large dip", []int{7, 0, 0, 0, 0, 7}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, trap(tt.height))
		})
	}
}

// Optional benchmark with a realistic large input.
func Benchmark_trap(b *testing.B) {
	h := make([]int, 10000)
	for i := range h {
		h[i] = i % 512
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trap(h)
	}
}
