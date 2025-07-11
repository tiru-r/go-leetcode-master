package maximum_product_subarray_152

// Note: study again.
// Time: O(n), Space: O(1)
func maxProduct(nums []int) int {
	minCurrent, maxCurrent, result := nums[0], nums[0], nums[0]

	for _, num := range nums[1:] {
		if num < 0 {
			minCurrent, maxCurrent = maxCurrent, minCurrent
		}

		minCurrent = min(num, minCurrent*num)
		maxCurrent = max(num, maxCurrent*num)
		result = max(result, maxCurrent)
	}

	return result
}

// First, O(n^2) solution.
func maxProduct0(nums []int) int {
	result := nums[0]
	for i, num := range nums {
		product := num
		result = max(result, product)
		for _, nextNum := range nums[i+1:] {
			product *= nextNum
			result = max(result, product)
		}
	}

	return result
}
