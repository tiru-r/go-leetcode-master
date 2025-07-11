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

// REMOVED: O(n²) brute force solution
// This approach checked all possible subarrays with nested loops,
// resulting in O(n²) time complexity. Use the O(n) Kadane variant above instead.
