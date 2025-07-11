package product_of_array_except_self_238

// Ultra-efficient O(n) time, O(1) extra space using modern range patterns
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate left products with modern range-over-int
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Calculate right products and combine - reverse iteration
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}
