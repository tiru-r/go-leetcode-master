package maximum_average_subarray_i_643

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}
