package moving_average_from_data_stream_346

import "slices"

// Modern circular buffer using slice with optimized operations
type MovingAverage struct {
	data  []int
	head  int
	size  int
	sum   int
	count int
}

// Legacy slice-based implementation for comparison
type MovingAverageSlice struct {
	q    []int
	size int
	sum  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	data := make([]int, size)
	data = slices.Grow(data, size)
	return MovingAverage{
		data:  data[:size],
		head:  0,
		size:  size,
		sum:   0,
		count: 0,
	}
}

// Alternative constructor for slice-based version
func ConstructorSlice(size int) MovingAverageSlice {
	return MovingAverageSlice{
		q:    make([]int, 0, size), // Pre-allocate capacity
		size: size,
		sum:  0,
	}
}

func (ma *MovingAverage) Next(val int) float64 {
	if ma.count == ma.size {
		// Buffer is full, subtract the value we're overwriting
		ma.sum -= ma.data[ma.head]
	} else {
		ma.count++
	}

	// Set new value at current head position
	ma.data[ma.head] = val
	ma.sum += val

	// Move head to next position (circular)
	ma.head = (ma.head + 1) % ma.size

	return float64(ma.sum) / float64(ma.count)
}

// Alternative slice-based implementation with modern Go patterns
func (mas *MovingAverageSlice) Next(val int) float64 {
	if len(mas.q) == mas.size {
		// Remove first element
		mas.sum -= mas.q[0]
		mas.q = mas.q[1:]
	}

	mas.q = append(mas.q, val)
	mas.sum += val

	return float64(mas.sum) / float64(len(mas.q))
}
