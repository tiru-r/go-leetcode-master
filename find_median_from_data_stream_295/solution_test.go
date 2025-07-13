package find_median_from_data_stream_295

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []float64
	}{
		{"empty", []int{}, []float64{}},
		{"single", []int{1}, []float64{1}},
		{"two", []int{1, 2}, []float64{1, 1.5}},
		{"three", []int{1, 2, 3}, []float64{1, 1.5, 2}},
		{"descending", []int{5, 4, 3, 2, 1}, []float64{5, 4.5, 4, 3.5, 3}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mf := NewMedianFinder()
			for i, n := range tc.nums {
				mf.AddNum(n)
				got := mf.FindMedian()
				if math.Abs(got-tc.want[i]) > 1e-9 {
					t.Fatalf("after %v: want %.1f, got %.1f", tc.nums[:i+1], tc.want[i], got)
				}
			}
		})
	}
}

func BenchmarkAddNum(b *testing.B) {
	mf := NewMedianFinder()
	for i := 0; i < b.N; i++ {
		mf.AddNum(i)
	}
}
