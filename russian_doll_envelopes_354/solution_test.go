package russian_doll_envelopes_354

import "testing"

func TestMaxEnvelopes(t *testing.T) {
	tests := []struct {
		envelopes [][]int
		expected  int
	}{
		{[][]int{}, 0},
		{[][]int{{5, 4}}, 1},
		{[][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][]int{{1, 1}, {1, 1}, {1, 1}}, 1},
		{[][]int{{2, 1}, {3, 2}, {4, 3}, {5, 4}}, 4},
		{[][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}, 4},
		{[][]int{{1, 3}, {3, 5}, {6, 7}, {6, 8}, {8, 4}, {9, 5}}, 3},
	}

	for i, test := range tests {
		result := MaxEnvelopes(test.envelopes)
		if result != test.expected {
			t.Errorf("Test %d: MaxEnvelopes(%v) = %d; want %d", i, test.envelopes, result, test.expected)
		}
	}
}

func BenchmarkMaxEnvelopes(b *testing.B) {
	envelopes := make([][]int, 1000)
	for i := range envelopes {
		envelopes[i] = []int{i + 1, i + 1}
	}

	b.ResetTimer()
	for range b.N {
		MaxEnvelopes(envelopes)
	}
}
