package maximum_subarray_53

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currSum := nums[0]

	for _, num := range nums[1:] {
		currSum = max(num, currSum+num)
		maxSum = max(maxSum, currSum)
	}

	return maxSum
}
