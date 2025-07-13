package russian_doll_envelopes_354

import (
	"slices"
	"sort"
)

// MaxEnvelopes finds the maximum number of envelopes that can be Russian dolled.
// Time: O(n log n), Space: O(n)
func MaxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	
	// Sort by width ascending, height descending for same width
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	
	// Find LIS on heights using binary search
	var tails []int
	for _, env := range envelopes {
		height := env[1]
		pos, _ := slices.BinarySearch(tails, height)
		if pos == len(tails) {
			tails = append(tails, height)
		} else {
			tails[pos] = height
		}
	}
	
	return len(tails)
}
