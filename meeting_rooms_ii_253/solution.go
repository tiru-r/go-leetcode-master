package meeting_rooms_ii_253

import (
	"cmp"
	"slices"
)

// minMeetingRooms returns the minimum number of conference rooms required
// to accommodate all meetings without overlap.
// Time: O(n log n) due to sorting, Space: O(n) for start and end time arrays.
func minMeetingRooms(intervals [][]int) int {
	// Handle empty input
	if len(intervals) == 0 {
		return 0
	}

	// Extract start and end times
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))
	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	// Sort start times and end times independently
	slices.SortFunc(starts, func(a, b int) int { return cmp.Compare(a, b) })
	slices.SortFunc(ends, func(a, b int) int { return cmp.Compare(a, b) })

	// Track maximum rooms needed and current end index
	maxRooms := 0
	currentRooms := 0
	endIdx := 0

	// Process each start time
	for _, start := range starts {
		// Free up rooms for meetings that have ended
		for endIdx < len(ends) && ends[endIdx] <= start {
			currentRooms--
			endIdx++
		}
		// Allocate a room for the new meeting
		currentRooms++
		// Update maximum rooms needed
		if currentRooms > maxRooms {
			maxRooms = currentRooms
		}
	}

	return maxRooms
}
