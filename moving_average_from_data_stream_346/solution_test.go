package moving_average_from_data_stream_346

import (
	"testing"
)

func TestMovingAverage(t *testing.T) {
	tests := []struct {
		size     int
		values   []int
		expected []float64
	}{
		{3, []int{1, 10, 3, 5}, []float64{1, 5.5, 14.0 / 3, 6}},
		{1, []int{4, 0, 3, 1}, []float64{4, 0, 3, 1}},
		{2, []int{}, []float64{}},
	}

	for _, tt := range tests {
		ma := New(tt.size)
		var results []float64
		for _, value := range tt.values {
			results = append(results, ma.Next(value))
		}
		if !floatSliceEq(results, tt.expected, 1e-9) {
			t.Errorf("size %d: got %v, want %v", tt.size, results, tt.expected)
		}
	}
}

func BenchmarkNext(b *testing.B) {
	ma := New(1000)
	b.ResetTimer()
	for i := range b.N {
		_ = ma.Next(i)
	}
}

func floatSliceEq(a, b []float64, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] < b[i]-eps || a[i] > b[i]+eps {
			return false
		}
	}
	return true
}
