package product_of_array_except_self_238

// Uses prefix and suffix products to calculate the answer.
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate left products
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Calculate right products and combine with left products
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return result
}
