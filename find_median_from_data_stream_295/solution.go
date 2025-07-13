package find_median_from_data_stream_295

import "container/heap"

type intHeap struct {
	data []int
	less func(a, b int) bool
}

func (h intHeap) Len() int           { return len(h.data) }
func (h intHeap) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h intHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *intHeap) Push(x any)        { h.data = append(h.data, x.(int)) }
func (h *intHeap) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}

type MedianFinder struct {
	left  intHeap // max heap for smaller half
	right intHeap // min heap for larger half
}

func NewMedianFinder() *MedianFinder {
	return &MedianFinder{
		left:  intHeap{less: func(a, b int) bool { return a > b }}, // max heap
		right: intHeap{less: func(a, b int) bool { return a < b }}, // min heap
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.left.Len() == 0 || num <= mf.left.data[0] {
		heap.Push(&mf.left, num)
	} else {
		heap.Push(&mf.right, num)
	}
	
	// Balance heaps
	if mf.left.Len() > mf.right.Len()+1 {
		heap.Push(&mf.right, heap.Pop(&mf.left))
	} else if mf.right.Len() > mf.left.Len()+1 {
		heap.Push(&mf.left, heap.Pop(&mf.right))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.left.Len() > mf.right.Len() {
		return float64(mf.left.data[0])
	}
	if mf.right.Len() > mf.left.Len() {
		return float64(mf.right.data[0])
	}
	return float64(mf.left.data[0]+mf.right.data[0]) / 2.0
}
