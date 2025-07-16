package russian_doll_envelopes_354

import "slices"

func MaxEnvelopes(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	
	slices.SortFunc(envelopes, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return b[1] - a[1]
	})
	
	var tails []int
	for _, env := range envelopes {
		pos, _ := slices.BinarySearch(tails, env[1])
		if pos == len(tails) {
			tails = append(tails, env[1])
		} else {
			tails[pos] = env[1]
		}
	}
	
	return len(tails)
}
