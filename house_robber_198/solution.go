// Package house_robber_198 solves LeetCode 198.
package house_robber_198

// rob returns the maximum money that can be robbed without alerting the police.
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

	prev2, prev1 := 0, nums[0]
	for i := 1; i < n; i++ {
		curr := max(prev1, nums[i]+prev2)
		prev2, prev1 = prev1, curr
	}
	return prev1
}
