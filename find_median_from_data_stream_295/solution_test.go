package find_median_from_data_stream_295

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		adds []int
		want []float64
	}{
		{[]int{1, 2, 3}, []float64{1.0, 1.5, 2.0}},
		{[]int{-1, -2, -3, -4, -5}, []float64{-1.0, -1.5, -2.0, -2.5, -3.0}},
		{[]int{6, 10, 2, 6, 5, 0, 6, 3, 1, 0, 0}, []float64{6.0, 8.0, 6.0, 6.0, 6.0, 5.5, 6.0, 5.5, 5.0, 5.0, 5.0}},
		{[]int{2, 3, 4}, []float64{2.0, 2.5, 3.0}},
		{[]int{}, []float64{}},
		{[]int{42}, []float64{42.0}},
	}

	for _, tt := range tests {
		mf := Constructor()
		for i, v := range tt.adds {
			mf.AddNum(v)
			assert.Equal(t, tt.want[i], mf.FindMedian())
		}
	}
}
