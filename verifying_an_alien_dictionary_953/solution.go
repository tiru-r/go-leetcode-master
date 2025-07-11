package verifying_an_alien_dictionary_953

import (
	"cmp"
	"slices"
)

func isAlienSorted(words []string, order string) bool {
	// Use byte keys for better performance than string keys
	charToIndex := make(map[byte]int)
	for idx, ch := range order {
		charToIndex[byte(ch)] = idx
	}

	// Use slices.IsSortedFunc with custom alien comparison
	return slices.IsSortedFunc(words, func(a, b string) int {
		return compareAlienStrings(a, b, charToIndex)
	})
}

// Returns -1 if s1 < s2, 0 if s1 == s2, and 1 if s1 > s2 using modern cmp
func compareAlienStrings(s1, s2 string, charToIndex map[byte]int) int {
	if s1 == s2 {
		return 0
	}

	for i := 0; i < min(len(s1), len(s2)); i++ {
		// Use direct byte access for better performance
		idx1 := charToIndex[s1[i]]
		idx2 := charToIndex[s2[i]]

		if result := cmp.Compare(idx1, idx2); result != 0 {
			return result
		}
	}

	// Compare lengths using cmp.Compare
	return cmp.Compare(len(s1), len(s2))
}
