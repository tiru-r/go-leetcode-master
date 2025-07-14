package house_robber_ii_213

// rob returns maximum money from circular street using optimized dual-pass DP
// Optimized: O(n) time, O(1) space with Go 1.24 modern syntax
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	
	// Circular constraint: first and last house are adjacent
	// Solution: max(rob houses 0..n-2, rob houses 1..n-1)
	return max(robLinear(nums[:n-1]), robLinear(nums[1:]))
}

// robLinear solves linear house robber with space-efficient DP
func robLinear(houses []int) int {
	rob, skip := 0, 0
	for _, money := range houses {
		rob, skip = max(skip+money, rob), rob
	}
	return rob
}
