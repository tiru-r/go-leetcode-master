package top_k_frequent_elements_347

import (
	"cmp"
	"slices"
)

// Modern Go 1.24 solution using slices.SortFunc - more efficient than heap for this use case
func topKFrequent(nums []int, k int) []int {
	// Build frequency map
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// Get unique numbers and sort by frequency using modern slices operations
	unique := make([]int, 0, len(freq))
	for key := range freq {
		unique = append(unique, key)
	}
	
	// Use Go 1.24 slices.SortFunc with cmp.Compare for type safety
	slices.SortFunc(unique, func(a, b int) int {
		return cmp.Compare(freq[b], freq[a]) // Descending order by frequency
	})

	return unique[:k]
}

// Alternative implementation using slices.SortFunc
func topKFrequentSlices(nums []int, k int) []int {
	// Build frequency map
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	// Get unique numbers and sort by frequency
	unique := make([]int, 0, len(freq))
	for k := range freq {
		unique = append(unique, k)
	}
	
	slices.SortFunc(unique, func(a, b int) int {
		return cmp.Compare(freq[b], freq[a]) // Descending order
	})

	return unique[:k]
}
