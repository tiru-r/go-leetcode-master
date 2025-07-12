package insert_interval_57

import "slices"

// Note: cool hard problem :)
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	var merged [][]int

	// discard all non-overlapping intervals by placing them into merged
	i := 0
	for i < len(intervals) {
		if intervals[i][1] < newInterval[0] {
			merged = append(merged, intervals[i])
		} else {
			break
		}

		i++
	}

	// 0: did not find an interval where new interval start is less than i-th interval end
	if i == len(intervals) {
		return append(merged, newInterval)
	}

	// 1: interval doesn't merge at all
	if newInterval[0] < intervals[i][0] && newInterval[1] < intervals[i][0] {
		return slices.Concat(merged, [][]int{newInterval}, intervals[i:])
	}

	// the minimum merged value is the new interval start or the i-th interval start
	minMerged := min(newInterval[0], intervals[i][0])

	// 2: the new interval end is between the i-th start and end inclusive
	if newInterval[1] >= intervals[i][0] && newInterval[1] <= intervals[i][1] {
		maxMerged := intervals[i][1]
		merged = append(merged, []int{minMerged, maxMerged})
		i++
	} else {
		// 3: the new interval end could span many intervals
		maxMerged := -1
		i++
		for i < len(intervals) {
			if newInterval[1] < intervals[i][0] {
				maxMerged = newInterval[1]
				break
			} else if newInterval[1] >= intervals[i][0] && newInterval[1] <= intervals[i][1] {
				// if new interval end falls in between an interval
				maxMerged = intervals[i][1]
				i++
				break
			}

			i++
		}

		// new interval end is larger than all other intervals
		if maxMerged == -1 {
			maxMerged = newInterval[1]
		}

		merged = append(merged, []int{minMerged, maxMerged})
	}

	// place the tail (n + 1) intervals into merged
	for i < len(intervals) {
		merged = append(merged, intervals[i])
		i++
	}

	return merged
}

