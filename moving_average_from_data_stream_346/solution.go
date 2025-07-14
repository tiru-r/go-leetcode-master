package moving_average_from_data_stream_346

type MovingAverage struct {
	buffer []int
	idx    int
	size   int
	sum    int
	count  int
}

func New(size int) *MovingAverage {
	return &MovingAverage{
		buffer: make([]int, size),
		size:   size,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	if m.count < m.size {
		m.count++
	} else {
		m.sum -= m.buffer[m.idx]
	}

	m.buffer[m.idx] = val
	m.sum += val
	m.idx = (m.idx + 1) % m.size

	return float64(m.sum) / float64(m.count)
}
