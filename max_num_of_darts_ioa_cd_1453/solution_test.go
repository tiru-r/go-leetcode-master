package maximum_number_of_darts_inside_of_a_circular_dartboard_1453

import (
	"math/rand"
	"testing"
)

func Test_numPointsInCircle(t *testing.T) {
	tests := []struct {
		name  string
		darts [][2]float64
		r     float64
		want  int
	}{
		{
			name:  "four points in cross pattern",
			darts: [][2]float64{{-2, 0}, {2, 0}, {0, 2}, {0, -2}},
			r:     2.0,
			want:  4,
		},
		{
			name:  "complex case",
			darts: [][2]float64{{-3, 0}, {3, 0}, {2, 6}, {5, 4}, {0, 9}, {7, 8}},
			r:     5.0,
			want:  5,
		},
		{
			name:  "single point",
			darts: [][2]float64{{0, 0}},
			r:     1.0,
			want:  1,
		},
		{
			name:  "two adjacent points",
			darts: [][2]float64{{0, 0}, {1, 0}},
			r:     1.0,
			want:  2,
		},
		{
			name:  "two distant points",
			darts: [][2]float64{{0, 0}, {3, 0}},
			r:     1.0,
			want:  1,
		},
		{
			name:  "clustered points",
			darts: [][2]float64{{1, 2}, {3, 5}, {1, -1}, {2, 3}, {4, 1}, {1, 3}},
			r:     2.0,
			want:  4,
		},
		{
			name:  "triangle points",
			darts: [][2]float64{{0, 0}, {0, 1}, {1, 0}},
			r:     1.0,
			want:  3,
		},
		{
			name:  "empty darts",
			darts: [][2]float64{},
			r:     1.0,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numPointsInCircle(tt.darts, tt.r)
			if got != tt.want {
				t.Errorf("numPointsInCircle(%v, %.1f) = %d; want %d", tt.darts, tt.r, got, tt.want)
			}
		})
	}
}

// Helper function tests removed as they test internal implementation details

func Benchmark_numPointsInCircle(b *testing.B) {
	darts := [][2]float64{{-2, 0}, {2, 0}, {0, 2}, {0, -2}, {1, 1}, {-1, -1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numPointsInCircle(darts, 2.0)
	}
}

func Benchmark_numPointsInCircleLarge(b *testing.B) {
	// Generate random darts
	darts := make([][2]float64, 50)
	for i := range darts {
		darts[i] = [2]float64{float64(rand.Intn(200) - 100), float64(rand.Intn(200) - 100)}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numPointsInCircle(darts, 10.0)
	}
}
