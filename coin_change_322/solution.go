package coin_change_322

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	
	dp := make([]int, amount+1)
	for i := range dp[1:] {
		dp[i+1] = amount + 1
	}
	
	for a := range amount {
		a++
		for _, coin := range coins {
			if coin <= a {
				dp[a] = min(dp[a], dp[a-coin]+1)
			}
		}
	}
	
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
