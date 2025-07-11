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

func (this *MovingAverage) Next(val int) float64 {
	if this.count == this.size {
		// Buffer is full, subtract the value we're overwriting
		this.sum -= this.data[this.head]
	} else {
		this.count++
	}
	
	// Set new value at current head position
	this.data[this.head] = val
	this.sum += val
	
	// Move head to next position (circular)
	this.head = (this.head + 1) % this.size
	
	return float64(this.sum) / float64(this.count)
}

// Alternative slice-based implementation with modern Go patterns
func (this *MovingAverageSlice) Next(val int) float64 {
	if len(this.q) == this.size {
		// Remove first element
		this.sum -= this.q[0]
		this.q = this.q[1:]
	}
	
	this.q = append(this.q, val)
	this.sum += val
	
	return float64(this.sum) / float64(len(this.q))
}
