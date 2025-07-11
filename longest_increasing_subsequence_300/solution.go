package longest_increasing_subsequence_300

import (
	"slices"
)

// Modern O(n log n) solution using binary search with slices.BinarySearch
// This is significantly faster than the O(n²) DP approach for large inputs
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] represents the smallest ending element of all increasing subsequences of length i+1
	tails := make([]int, 0, len(nums)/2) // Pre-allocate with estimated capacity

	for _, num := range nums {
		// Use Go 1.21+ slices.BinarySearch for efficient searching
		pos, found := slices.BinarySearch(tails, num)

		if found {
			// Skip duplicates to maintain strictly increasing property
			continue
		}

		if pos == len(tails) {
			// num is larger than all elements in tails, extend the sequence
			tails = append(tails, num)
		} else {
			// Replace the element at pos to maintain smallest possible ending values
			tails[pos] = num
		}
	}

	return len(tails)
}
