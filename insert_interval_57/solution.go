package insert_interval_57

// insert inserts newInterval into intervals, merging where necessary.
// Time: O(n)  Space: O(1) (output slice excluded)
func insert(intervals [][]int, newInterval []int) [][]int {
	// short-circuit empty input
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var out [][]int
	i, start, end := 0, newInterval[0], newInterval[1]

	// 1. Append intervals completely before newInterval
	for i < len(intervals) && intervals[i][1] < start {
		out = append(out, intervals[i])
		i++
	}

	// 2. Merge overlapping / touching intervals
	for i < len(intervals) && intervals[i][0] <= end {
		start = min(start, intervals[i][0])
		end = max(end, intervals[i][1])
		i++
	}
	out = append(out, []int{start, end})

	// 3. Append remaining intervals
	out = append(out, intervals[i:]...)

	return out
}
