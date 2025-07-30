package the_skyline_problem_218

import (
	"container/heap"
	"sort"
)

// Event represents a building boundary event in the sweep line algorithm.
type Event struct {
	X      int  // x-coordinate
	Height int  // building height
	IsEnd  bool // true if this is a building end event
}

// MaxHeap implements a max-heap for building heights.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// GetSkyline returns the skyline formed by the given buildings.
// Time: O(n log n), Space: O(n)
func GetSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	// Create events for building starts and ends
	events := make([]Event, 0, 2*len(buildings))
	for _, building := range buildings {
		left, right, height := building[0], building[1], building[2]
		events = append(events, Event{X: left, Height: height, IsEnd: false})
		events = append(events, Event{X: right, Height: height, IsEnd: true})
	}

	// Sort events by x-coordinate with proper tie-breaking
	sort.Slice(events, func(i, j int) bool {
		if events[i].X != events[j].X {
			return events[i].X < events[j].X
		}
		// Same x-coordinate: process building starts before ends
		if events[i].IsEnd != events[j].IsEnd {
			return !events[i].IsEnd
		}
		// Both starts: process higher buildings first
		if !events[i].IsEnd {
			return events[i].Height > events[j].Height
		}
		// Both ends: process lower buildings first
		return events[i].Height < events[j].Height
	})

	// Process events using sweep line algorithm with max-heap
	heights := &MaxHeap{0} // Start with ground level
	heap.Init(heights)
	prevHeight := 0
	toDelete := make(map[int]int) // Lazy deletion map

	result := make([][]int, 0)

	for _, event := range events {
		if !event.IsEnd {
			// Building start: add height to heap
			heap.Push(heights, event.Height)
		} else {
			// Building end: mark height for lazy deletion
			toDelete[event.Height]++
		}

		// Clean up deleted heights from top of heap
		for heights.Len() > 0 {
			top := (*heights)[0]
			if count, exists := toDelete[top]; exists && count > 0 {
				toDelete[top]--
				if toDelete[top] == 0 {
					delete(toDelete, top)
				}
				heap.Pop(heights)
			} else {
				break
			}
		}

		// Check if height changed
		currentHeight := (*heights)[0]
		if currentHeight != prevHeight {
			result = append(result, []int{event.X, currentHeight})
			prevHeight = currentHeight
		}
	}

	return result
}
