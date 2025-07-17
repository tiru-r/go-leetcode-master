package verifying_an_alien_dictionary_953

import (
	"cmp"
	"slices"
)

func isAlienSorted(words []string, order string) bool {
	charToIndex := make(map[byte]int)
	for idx, ch := range order {
		charToIndex[byte(ch)] = idx
	}

	return slices.IsSortedFunc(words, func(a, b string) int {
		return compareAlienStrings(a, b, charToIndex)
	})
}

func compareAlienStrings(s1, s2 string, charToIndex map[byte]int) int {
	if s1 == s2 {
		return 0
	}

	for i := range min(len(s1), len(s2)) {
		if result := cmp.Compare(charToIndex[s1[i]], charToIndex[s2[i]]); result != 0 {
			return result
		}
	}

	return cmp.Compare(len(s1), len(s2))
}
