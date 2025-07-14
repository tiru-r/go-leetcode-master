package longest_increasing_subsequence_300

import "slices"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Optimized binary search approach with patience sorting
	tails := make([]int, 0, len(nums))

	for _, num := range nums {
		// Find the leftmost position where tails[pos] >= num
		pos, found := slices.BinarySearch(tails, num)
		
		if !found {
			if pos == len(tails) {
				// Extend the sequence
				tails = append(tails, num)
			} else {
				// Replace with smaller ending value
				tails[pos] = num
			}
		}
		// If found, we skip duplicates to maintain strictly increasing
	}
	
	return len(tails)
}
