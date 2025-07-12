package longest_increasing_subsequence_300

import "slices"

// lengthOfLIS returns the length of the longest strictly increasing subsequence.
// Time: O(n log n)   Space: O(n)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := make([]int, 0, len(nums)) // final capacity = max possible LIS length

	for _, num := range nums {
		// pos = first index where tails[pos] >= num
		pos, _ := slices.BinarySearch(tails, num)

		if pos == len(tails) {
			tails = append(tails, num) // extend LIS
		} else {
			tails[pos] = num // shrink ending value
		}
	}
	return len(tails)
}
