package missing_number_268

func missingNumber(nums []int) int {
	n := len(nums)
	result := n

	for i, num := range nums {
		result ^= i ^ num
	}

	return result
}
