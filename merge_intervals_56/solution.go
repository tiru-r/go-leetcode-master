package merge_intervals_56

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]
		if interval[0] <= last[1] {
			last[1] = max(last[1], interval[1])
		} else {
			merged = append(merged, interval)
		}
	}

	return merged
}
