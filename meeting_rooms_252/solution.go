package meeting_rooms_252

import (
	"cmp"
	"slices"
)

func canAttendMeetings(intervals [][]int) bool {
	// Sort the intervals by start time using cmp.Compare for type safety
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// If any intervals overlap, one person could not attend all meetings
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}
