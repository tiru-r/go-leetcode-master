package group_anagrams_49

import (
	"maps"
	"slices"
)

// GroupAnagrams groups all anagrams together using optimized Go 1.24 features
// Optimized: O(N·L) time, O(N) space with character frequency grouping
func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	// Use character frequency array as map key for O(L) grouping per string
	groups := make(map[[26]byte][]string, len(strs)>>1+1)

	// Reusable frequency key to avoid allocations
	var key [26]byte
	for _, s := range strs {
		clear(key[:])
		
		// Count character frequencies using range over string (Go 1.24)
		for _, char := range []byte(s) {
			key[char-'a']++
		}
		
		groups[key] = append(groups[key], s)
	}

	// Extract groups using modern Go 1.24 slices.Collect with maps.Values
	return slices.Collect(maps.Values(groups))
}
