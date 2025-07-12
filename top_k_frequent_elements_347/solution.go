package top_k_frequent_elements_347

import (
	"cmp"
	"maps"
	"slices"
)

// Ultra-modern Go 1.24+ solution using functional programming patterns
func topKFrequent(nums []int, k int) []int {
	// Build frequency map with better initial capacity
	freq := make(map[int]int, len(nums)/2)
	for _, n := range nums {
		freq[n]++
	}

	// Get unique numbers using modern maps.Keys()
	unique := slices.Collect(maps.Keys(freq))
	
	// Sort by frequency in descending order
	slices.SortFunc(unique, func(a, b int) int {
		return cmp.Compare(freq[b], freq[a]) // Descending by frequency
	})

	return unique[:k]
}

