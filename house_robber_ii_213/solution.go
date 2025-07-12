package house_robber_ii_213

// rob returns the maximum money that can be robbed from a circular street.
// Time: O(n)  Space: O(1)
func rob(nums []int) int {
	n := len(nums)
	switch n {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}

	// Two linear scans:
	// 1. skip last house  2. skip first house
	return max(robLinear(nums[0:n-1]), robLinear(nums[1:]))
}

// robLinear solves the linear version (LeetCode 198).
func robLinear(nums []int) int {
	prev2, prev1 := 0, 0
	for _, v := range nums {
		prev2, prev1 = prev1, max(prev1, v+prev2)
	}
	return prev1
}
