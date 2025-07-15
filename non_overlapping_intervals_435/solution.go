package non_overlapping_intervals_435

import (
	"cmp"
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n <= 1 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < n; i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}
