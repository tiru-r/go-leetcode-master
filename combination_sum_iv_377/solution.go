package combination_sum_iv_377

func combinationSum4(nums []int, target int) int {
	if target == 0 {
		return 1
	}
	if target < 0 {
		return 0
	}
	
	dp := make([]int, target+1)
	dp[0] = 1
	
	for i := range target {
		i++
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	
	return dp[target]
}
