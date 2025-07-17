package top_k_frequent_elements_347

import (
	"cmp"
	"maps"
	"slices"
)

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums)/2)
	for _, n := range nums {
		freq[n]++
	}

	unique := slices.Collect(maps.Keys(freq))
	slices.SortFunc(unique, func(a, b int) int {
		return cmp.Compare(freq[b], freq[a])
	})

	return unique[:k]
}
