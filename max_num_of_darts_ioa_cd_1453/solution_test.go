package maximum_number_of_darts_inside_of_a_circular_dartboard_1453

import (
	"math"
	"testing"
)

func TestNumPointsInCircle(t *testing.T) {
	tests := []struct {
		darts [][]int
		r     int
		want  int
	}{
		{
			darts: [][]int{{-2, 0}, {2, 0}, {0, 2}, {0, -2}},
			r:     2,
			want:  4,
		},
		{
			darts: [][]int{{-3, 0}, {3, 0}, {2, 6}, {5, 4}, {0, 9}, {7, 8}},
			r:     5,
			want:  5,
		},
		{
			darts: [][]int{{0, 0}},
			r:     1,
			want:  1,
		},
		{
			darts: [][]int{{0, 0}, {1, 0}},
			r:     1,
			want:  2,
		},
		{
			darts: [][]int{{0, 0}, {3, 0}},
			r:     1,
			want:  1,
		},
		{
			darts: [][]int{{1, 2}, {3, 5}, {1, -1}, {2, 3}, {4, 1}, {1, 3}},
			r:     2,
			want:  4,
		},
		{
			darts: [][]int{{0, 0}, {0, 1}, {1, 0}},
			r:     1,
			want:  3,
		},
	}

	for i, test := range tests {
		got := NumPointsInCircle(test.darts, test.r)
		if got != test.want {
			t.Errorf("Test %d: NumPointsInCircle(%v, %d) = %d; want %d", i, test.darts, test.r, got, test.want)
		}
	}
}

func TestGetCircleCenters(t *testing.T) {
	// Test basic circle center finding
	p1 := Point{0, 0}
	p2 := Point{2, 0}
	centers := getCircleCenters(p1, p2, math.Sqrt(2))
	
	if len(centers) != 2 {
		t.Errorf("Expected 2 centers, got %d", len(centers))
	}
}

func TestDistance(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{3, 4}
	expected := 5.0
	
	got := distance(p1, p2)
	if math.Abs(got-expected) > EPS {
		t.Errorf("distance(%v, %v) = %f; want %f", p1, p2, got, expected)
	}
}

func BenchmarkNumPointsInCircle(b *testing.B) {
	darts := [][]int{{-2, 0}, {2, 0}, {0, 2}, {0, -2}, {1, 1}, {-1, -1}}
	
	b.ResetTimer()
	for range b.N {
		NumPointsInCircle(darts, 2)
	}
}
