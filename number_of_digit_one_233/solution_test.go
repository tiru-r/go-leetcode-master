package number_of_digit_one_233

import "testing"

func TestCountDigitOne(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{9, 1},
		{10, 2},
		{13, 6},
		{99, 20},
		{100, 21},
		{824, 273},
		{1000, 301},
		{9999, 4000},
	}

	for _, test := range tests {
		result := CountDigitOne(test.n)
		if result != test.expected {
			t.Errorf("CountDigitOne(%d) = %d; want %d", test.n, result, test.expected)
		}
	}
}

func BenchmarkCountDigitOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountDigitOne(999999999)
	}
}
