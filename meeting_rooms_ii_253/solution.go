package meeting_rooms_ii_253

import "slices"

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))
	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	slices.Sort(starts)
	slices.Sort(ends)

	maxRooms := 0
	currentRooms := 0
	endIdx := 0

	for _, start := range starts {
		for endIdx < len(ends) && ends[endIdx] <= start {
			currentRooms--
			endIdx++
		}
		currentRooms++
		maxRooms = max(maxRooms, currentRooms)
	}

	return maxRooms
}
