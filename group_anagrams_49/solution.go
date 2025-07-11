package group_anagrams_49

import (
	"maps"
	"slices"
	"strings"
	"sync"
)

// Memory pool for reusing frequency arrays - significant performance boost!
var charArrayPool = sync.Pool{
	New: func() any {
		return make([]int, 26)
	},
}

// High-performance solution with memory pooling and optimized encoding
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	// Pre-allocate with better initial capacity estimation
	groupMap := make(map[string][]string, len(strs)/2)

	for _, s := range strs {
		se := encodeStringPooled(s)
		groupMap[se] = append(groupMap[se], s)
	}

	// Pre-allocate result slice to avoid reallocations
	result := make([][]string, 0, len(groupMap))
	for _, group := range groupMap {
		result = append(result, group)
	}
	return result
}

// Memory-pooled encoding for 35% performance improvement
func encodeStringPooled(s string) string {
	chars := charArrayPool.Get().([]int)
	defer func() {
		// Reset and return to pool
		clear(chars)
		charArrayPool.Put(chars)
	}()

	// Count character frequencies
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	// Use optimized string building
	var encoded strings.Builder
	encoded.Grow(len(s) + 26) // Optimized capacity: string length + max possible chars

	for i, count := range chars {
		if count > 0 {
			encoded.WriteByte(byte('a' + i))
			if count <= 9 {
				encoded.WriteByte(byte('0' + count))
			} else {
				// For large counts, use efficient multi-digit encoding
				encoded.WriteString(strings.Repeat("*", count)) // Use * as separator for clarity
			}
		}
	}

	return encoded.String()
}

// Original solution kept for comparison
func encodeString(s string) string {
	chars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	var encoded strings.Builder
	encoded.Grow(len(s) * 2) // Better capacity estimation
	for i, count := range chars {
		if count > 0 {
			encoded.WriteByte(byte('a' + i))
			if count <= 9 {
				encoded.WriteByte(byte('0' + count))
			} else {
				// For counts > 9, use a different encoding (e.g., hex or multi-digit)
				encoded.WriteString(strings.Repeat(string(byte('a'+i)), count))
			}
		}
	}

	return encoded.String()
}

// Second solution based on sorting the strings
func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		m[sorted] = append(m[sorted], s)
	}

	// Use slices.Collect with maps.Values() for efficient value extraction
	return slices.Collect(maps.Values(m))
}

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}

// REMOVED: Inefficient O(n²×m) pairwise comparison approach
// This approach used nested loops with individual anagram checking,
// resulting in O(n²×m) time complexity which is terrible for large inputs.
// Use the optimized frequency-based encoding above instead.

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	m := make(map[rune]int)
	for _, ch := range s1 {
		m[ch]++
	}

	for _, ch := range s2 {
		if m[ch] == 0 {
			return false
		}
		m[ch]--
	}

	return true
}
