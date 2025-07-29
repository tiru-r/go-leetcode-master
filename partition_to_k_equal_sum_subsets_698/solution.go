package partition_to_k_equal_sum_subsets_698

import "slices"

// canPartitionKSubsets returns true if nums can be split into k
// non-empty subsets each summing to the same value using optimized backtracking.
func canPartitionKSubsets(nums []int, k int) bool {
	total := calculateSum(nums)
	if total%k != 0 {
		return false
	}
	target := total / k

	// Sort in ascending order for better pruning
	slices.Sort(nums)
	maxValue := nums[len(nums)-1]
	if maxValue > target {
		return false // Largest element exceeds target
	}

	// Optimization: remove exact matches to reduce search space
	lastIndex := trimExactMatches(nums, target, &k)
	if k == 0 {
		return true // All buckets filled with exact matches
	}

	// Use backtracking to fill remaining k buckets
	buckets := make([]int, k)
	return backtrackFill(buckets, lastIndex, nums, target)
}

// calculateSum returns the total sum of all numbers in the array
func calculateSum(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

// trimExactMatches removes elements that exactly equal target and decreases k accordingly.
// This optimization reduces the search space by handling perfect matches upfront.
func trimExactMatches(nums []int, target int, remainingBuckets *int) int {
	lastIndex := len(nums) - 1
	// Remove exact matches from the end (largest elements first due to sorting)
	for lastIndex >= 0 && nums[lastIndex] == target {
		lastIndex--
		*remainingBuckets-- // One less bucket needed
	}
	return lastIndex
}

// backtrackFill attempts to place nums[0..currentIndex] into buckets using backtracking.
// Key optimizations: early termination and symmetry breaking for empty buckets.
func backtrackFill(buckets []int, currentIndex int, nums []int, target int) bool {
	// Base case: all numbers have been placed successfully
	if currentIndex < 0 {
		return true
	}
	
	currentNumber := nums[currentIndex]

	// Try placing current number in each bucket
	for bucketIndex := range buckets {
		currentBucketSum := buckets[bucketIndex]
		
		// Check if current number fits in this bucket
		if currentBucketSum+currentNumber <= target {
			// Place number in bucket
			buckets[bucketIndex] = currentBucketSum + currentNumber
			
			// Recursively try to place remaining numbers
			if backtrackFill(buckets, currentIndex-1, nums, target) {
				return true
			}
			
			// Backtrack: remove number from bucket
			buckets[bucketIndex] = currentBucketSum
		}
		
		// Pruning optimization: if bucket is empty, skip remaining empty buckets
		// (they are equivalent due to symmetry)
		if currentBucketSum == 0 {
			break
		}
	}
	return false
}
