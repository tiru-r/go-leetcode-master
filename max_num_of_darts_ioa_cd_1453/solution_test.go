package maximum_number_of_darts_inside_of_a_circular_dartboard_1453

import (
	"math"
	"math/rand"
	"testing"
)

func Test_numPointsInCircle(t *testing.T) {
	tests := []struct {
		name  string
		darts [][]int
		r     int
		want  int
	}{
		{
			name:  "four points in cross pattern",
			darts: [][]int{{-2, 0}, {2, 0}, {0, 2}, {0, -2}},
			r:     2,
			want:  4,
		},
		{
			name:  "complex case",
			darts: [][]int{{-3, 0}, {3, 0}, {2, 6}, {5, 4}, {0, 9}, {7, 8}},
			r:     5,
			want:  5,
		},
		{
			name:  "single point",
			darts: [][]int{{0, 0}},
			r:     1,
			want:  1,
		},
		{
			name:  "two adjacent points",
			darts: [][]int{{0, 0}, {1, 0}},
			r:     1,
			want:  2,
		},
		{
			name:  "two distant points",
			darts: [][]int{{0, 0}, {3, 0}},
			r:     1,
			want:  1,
		},
		{
			name:  "clustered points",
			darts: [][]int{{1, 2}, {3, 5}, {1, -1}, {2, 3}, {4, 1}, {1, 3}},
			r:     2,
			want:  4,
		},
		{
			name:  "triangle points",
			darts: [][]int{{0, 0}, {0, 1}, {1, 0}},
			r:     1,
			want:  3,
		},
		{
			name:  "empty darts",
			darts: [][]int{},
			r:     1,
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numPointsInCircle(tt.darts, tt.r)
			if got != tt.want {
				t.Errorf("numPointsInCircle(%v, %d) = %d; want %d", tt.darts, tt.r, got, tt.want)
			}
		})
	}
}

func Test_getCircleCenters(t *testing.T) {
	tests := []struct {
		name     string
		p1, p2   point
		r        float64
		expected int
	}{
		{"two centers", point{0, 0}, point{2, 0}, math.Sqrt(2), 2},
		{"diameter apart", point{0, 0}, point{2, 0}, 1, 1},
		{"too far apart", point{0, 0}, point{10, 0}, 1, 0},
		{"same point", point{0, 0}, point{0, 0}, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			centers := getCircleCenters(tt.p1, tt.p2, tt.r)
			if len(centers) != tt.expected {
				t.Errorf("getCircleCenters() returned %d centers; want %d", len(centers), tt.expected)
			}
		})
	}
}

func Test_distanceSquared(t *testing.T) {
	p1 := point{0, 0}
	p2 := point{3, 4}
	expected := 25.0

	got := distanceSquared(p1, p2)
	if math.Abs(got-expected) > eps {
		t.Errorf("distanceSquared(%v, %v) = %f; want %f", p1, p2, got, expected)
	}
}

func Benchmark_numPointsInCircle(b *testing.B) {
	darts := [][]int{{-2, 0}, {2, 0}, {0, 2}, {0, -2}, {1, 1}, {-1, -1}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numPointsInCircle(darts, 2)
	}
}

func Benchmark_numPointsInCircleLarge(b *testing.B) {
	// Generate random darts
	darts := make([][]int, 50)
	for i := range darts {
		darts[i] = []int{rand.Intn(200) - 100, rand.Intn(200) - 100}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		numPointsInCircle(darts, 10)
	}
}
