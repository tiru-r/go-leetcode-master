package meeting_rooms_252

import (
	"cmp"
	"slices"
)

func canAttendMeetings(intervals [][]int) bool {
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}
