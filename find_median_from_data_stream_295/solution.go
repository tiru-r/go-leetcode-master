package find_median_from_data_stream_295

import "container/heap"

type minHeap []int
type maxHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any          { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }

type MedianFinder struct {
	lower *maxHeap
	upper *minHeap
}

func NewMedianFinder() *MedianFinder {
	lower, upper := new(maxHeap), new(minHeap)
	heap.Init(lower)
	heap.Init(upper)
	return &MedianFinder{lower: lower, upper: upper}
}

func (mf *MedianFinder) AddNum(num int) {
	heap.Push(mf.lower, num)
	heap.Push(mf.upper, heap.Pop(mf.lower).(int))
	if mf.upper.Len() > mf.lower.Len() {
		heap.Push(mf.lower, heap.Pop(mf.upper).(int))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.lower.Len() > mf.upper.Len() {
		return float64((*mf.lower)[0])
	}
	return (float64((*mf.lower)[0]) + float64((*mf.upper)[0])) / 2.0
}
