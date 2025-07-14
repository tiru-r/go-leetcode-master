package maximum_product_subarray_152

func maxProduct(nums []int) int {
	minProd, maxProd, result := nums[0], nums[0], nums[0]

	for _, num := range nums[1:] {
		if num < 0 {
			minProd, maxProd = maxProd, minProd
		}

		minProd = min(num, minProd*num)
		maxProd = max(num, maxProd*num)
		result = max(result, maxProd)
	}

	return result
}
