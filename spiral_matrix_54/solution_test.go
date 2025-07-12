package spiral_matrix_54

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{
			name:   "3x3",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name:   "2x3",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}},
			want:   []int{1, 2, 3, 6, 5, 4},
		},
		{
			name:   "3x2",
			matrix: [][]int{{1, 2}, {3, 4}, {5, 6}},
			want:   []int{1, 2, 4, 6, 5, 3},
		},
		{
			name:   "single row",
			matrix: [][]int{{1, 2, 3, 4}},
			want:   []int{1, 2, 3, 4},
		},
		{
			name:   "single column",
			matrix: [][]int{{1}, {2}, {3}},
			want:   []int{1, 2, 3},
		},
		{
			name:   "single element",
			matrix: [][]int{{42}},
			want:   []int{42},
		},
		{
			name:   "empty",
			matrix: [][]int{},
			want:   nil,
		},
		{
			name:   "row with empty slice",
			matrix: [][]int{{}},
			want:   nil,
		},
		{
			name:   "2x4",
			matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			want:   []int{1, 2, 3, 4, 8, 7, 6, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, spiralOrder(tt.matrix))
		})
	}
}

// BenchmarkSpiralOrder measures throughput on a large 1000×1000 matrix.
func BenchmarkSpiralOrder(b *testing.B) {
	const n = 1000
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m[i][j] = i*n + j
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = spiralOrder(m)
	}
}
