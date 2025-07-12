package group_anagrams_49

// GroupAnagrams groups all anagrams together and returns them in any order.
// Time: O(N·L)  (N = #strings, L = average length)
// Space: O(N)   (plus pooled 26-int arrays)
func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	// Key → list of original strings
	groups := make(map[[26]byte][]string, len(strs)/2+1)

	// Re-usable 26-byte key; avoid sync.Pool for simplicity & zero-copy
	var key [26]byte
	for _, s := range strs {
		clear(key[:]) // zero the array
		for i := 0; i < len(s); i++ {
			key[s[i]-'a']++
		}
		groups[key] = append(groups[key], s)
	}

	// Convert map → slice
	out := make([][]string, 0, len(groups))
	for _, g := range groups {
		out = append(out, g)
	}
	return out
}
