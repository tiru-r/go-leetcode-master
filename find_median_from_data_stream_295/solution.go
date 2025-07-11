package find_median_from_data_stream_295

import (
	"container/heap"
	"cmp"
)

// IntHeap implements heap.Interface for integers
type IntHeap struct {
	data    []int
	isMax   bool
}

func (h IntHeap) Len() int           { return len(h.data) }
func (h IntHeap) Less(i, j int) bool {
	if h.isMax {
		return cmp.Compare(h.data[i], h.data[j]) > 0
	}
	return cmp.Compare(h.data[i], h.data[j]) < 0
}
func (h IntHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *IntHeap) Push(x any) {
	h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() int {
	if len(h.data) == 0 {
		return 0
	}
	return h.data[0]
}

type MedianFinder struct {
	// keep the smaller half of numbers in a max heap for O(1) access to middle
	smallerHalf *IntHeap

	// keep the larger half of numbers in a min heap for O(1) access to middle
	largerHalf *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	largerHalf := &IntHeap{data: make([]int, 0), isMax: false}  // min heap
	smallerHalf := &IntHeap{data: make([]int, 0), isMax: true}  // max heap
	heap.Init(largerHalf)
	heap.Init(smallerHalf)
	return MedianFinder{
		largerHalf:  largerHalf,
		smallerHalf: smallerHalf,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// push num onto the max heap for the smaller half of numbers
	heap.Push(this.smallerHalf, num)

	// to maintain balance between the two heaps, offer
	// largest from smaller half to the larger half
	val := heap.Pop(this.smallerHalf).(int)
	heap.Push(this.largerHalf, val)

	// if the larger half becomes larger in length than the smaller half,
	// we need to offer the lowest in the largest half to the smaller half
	if this.largerHalf.Len() > this.smallerHalf.Len() {
		val := heap.Pop(this.largerHalf).(int)
		heap.Push(this.smallerHalf, val)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.smallerHalf.Len() == this.largerHalf.Len() {
		// there is an even number of elements in the MedianFinder, so the median is the
		// average of largest of the smaller half and the smallest of the larger half
		smaller := this.smallerHalf.Peek()
		larger := this.largerHalf.Peek()
		return (float64(smaller) + float64(larger)) / 2
	} else {
		// there is an odd number of elements in the MedianFinder,
		// so the median is the largest of the smaller half
		val := this.smallerHalf.Peek()
		return float64(val)
	}
}
