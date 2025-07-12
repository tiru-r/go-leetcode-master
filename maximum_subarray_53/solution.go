package maximum_subarray_53

// Note: study again
// Kadane's Algorithm works for this problem
// https://www.youtube.com/watch?v=86CQq3pKSUw
// Remembering past sums (dynamic programming) reduces runtime.
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	maxCurrent := nums[0]

	for i := 1; i < len(nums); i++ {
		maxCurrent = max(nums[i], maxCurrent+nums[i])
		maxSum = max(maxSum, maxCurrent)
	}

	return maxSum
}

