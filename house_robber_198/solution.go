package house_robber_198

// Modern solution using built-in max function (Go 1.21+)

// Solved alone for second time. Better solution.
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev2, prev1 := 0, nums[0]
	
	for i := 1; i < len(nums); i++ {
		current := max(prev1, nums[i]+prev2)
		prev2, prev1 = prev1, current
	}

	return prev1
}

