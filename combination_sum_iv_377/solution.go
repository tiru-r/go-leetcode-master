package combination_sum_iv_377

import "slices"

// combinationSum4 calculates the number of combinations of numbers from nums that sum to target.
// Numbers can be used multiple times, and different permutations are counted as distinct combinations.
// Uses dynamic programming with optimizations for early termination and sorted input.
func combinationSum4(nums []int, target int) int {
	// Handle edge cases: empty input array returns 0 as no combinations are possible.
	if len(nums) == 0 {
		return 0
	}

	// Early termination: if all numbers in nums are larger than target, no combinations are possible.
	allLarger := true
	for _, n := range nums {
		if n <= target {
			allLarger = false
			break
		}
	}
	if allLarger {
		return 0
	}

	// Create a sorted copy of nums to enable early loop termination, reducing iterations.
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	slices.Sort(sortedNums)

	// Initialize DP array where dp[i] represents the number of combinations summing to i.
	dp := make([]int, target+1)
	dp[0] = 1 // Base case: one way to make sum 0 (empty combination).

	// Build combinations for each sum from 1 to target.
	for i := 1; i <= target; i++ {
		// Iterate through sorted numbers to find valid contributions to sum i.
		for _, n := range sortedNums {
			// Skip numbers larger than current sum i, as they cannot contribute.
			// Break early since sortedNums ensures all subsequent numbers are larger.
			if n > i {
				break
			}
			// Add the number of combinations for sum (i-n) to dp[i], as n can be the last number.
			// Example: for i=3, n=1, add dp[2] to dp[3] for combinations ending with 1.
			dp[i] += dp[i-n]
		}
	}

	// Return the number of combinations that sum to target.
	return dp[target]
}
