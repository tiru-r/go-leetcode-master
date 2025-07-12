package moving_average_from_data_stream_346

// MovingAverage keeps the last `size` values and exposes their mean.
type MovingAverage struct {
	data  []int // circular buffer
	head  int   // index of oldest element
	size  int   // capacity
	sum   int   // running sum
	count int   // current #elements ( ≤ size )
}

// Constructor allocates the circular buffer.
func Constructor(size int) MovingAverage {
	return MovingAverage{
		data:  make([]int, size),
		head:  0,
		size:  size,
		sum:   0,
		count: 0,
	}
}

// Next adds a value and returns the current moving average.
func (ma *MovingAverage) Next(val int) float64 {
	if ma.count == ma.size {
		ma.sum -= ma.data[ma.head]
	} else {
		ma.count++
	}

	ma.data[ma.head] = val
	ma.sum += val
	ma.head = (ma.head + 1) % ma.size

	return float64(ma.sum) / float64(ma.count)
}
