package insert_interval_57

// insert merges newInterval into sorted intervals using optimized three-phase approach
// Optimized: O(n) time, O(1) space with Go 1.24 modern syntax
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0, len(intervals)+1)
	i := 0
	start, end := newInterval[0], newInterval[1]

	// Phase 1: Add all intervals that end before newInterval starts
	for i < len(intervals) && intervals[i][1] < start {
		result = append(result, intervals[i])
		i++
	}

	// Phase 2: Merge all overlapping/touching intervals with newInterval
	for i < len(intervals) && intervals[i][0] <= end {
		start = min(start, intervals[i][0])
		end = max(end, intervals[i][1])
		i++
	}
	result = append(result, []int{start, end})

	// Phase 3: Add all remaining intervals
	result = append(result, intervals[i:]...)

	return result
}
