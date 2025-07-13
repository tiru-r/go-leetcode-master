package moving_average_from_data_stream_346

type MovingAverage struct {
	buffer   []int
	index    int
	capacity int
	sum      int
	length   int
}

func NewMovingAverage(size int) *MovingAverage {
	return &MovingAverage{
		buffer:   make([]int, size),
		capacity: size,
	}
}

func (ma *MovingAverage) Next(val int) float64 {
	if ma.length < ma.capacity {
		ma.length++
	} else {
		ma.sum -= ma.buffer[ma.index]
	}

	ma.buffer[ma.index] = val
	ma.sum += val
	ma.index = (ma.index + 1) % ma.capacity

	return float64(ma.sum) / float64(ma.length)
}
