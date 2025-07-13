package max_points_on_a_line_149

import "testing"

func TestMaxPointsOnLine(t *testing.T) {
	tests := []struct {
		points   [][]int
		expected int
	}{
		{
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected: 3,
		},
		{
			points:   [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			expected: 4,
		},
		{
			points:   [][]int{{0, 0}},
			expected: 1,
		},
		{
			points:   [][]int{{0, 0}, {1, 1}},
			expected: 2,
		},
		{
			points:   [][]int{{0, 0}, {0, 0}},
			expected: 2,
		},
		{
			points:   [][]int{{1, 1}, {1, 1}, {1, 1}},
			expected: 3,
		},
		{
			points:   [][]int{{0, 0}, {1, 1}, {0, 0}},
			expected: 3,
		},
		{
			points:   [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 6}},
			expected: 3,
		},
	}

	for i, test := range tests {
		result := maxPointsOnLine(test.points)
		if result != test.expected {
			t.Errorf("Test %d failed: expected %d, got %d", i+1, test.expected, result)
		}
	}
}
