package coin_change_322

// Version 3: Bottom-up approach using filled slice with one above amount,
// which is one above max in scenario with denomination 1 * amount.
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for a := 1; a <= amount; a++ {
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

// Version 2: Bottom-up approach using MaxInt32 for min comparison
func coinChangeBottomUp1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for a := 1; a <= amount; a++ {
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

// Version 1: Top-down approach using recursion
func coinChangeTopDown(coins []int, amount int) int {
	if len(coins) == 0 || amount < 1 {
		return -1
	}
	return minChange(coins, amount, make([]int, amount))
}

func minChange(coins []int, rem int, memo []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}
	if memo[rem-1] != 0 {
		return memo[rem-1]
	}

	result := rem + 1
	for _, coin := range coins {
		if change := minChange(coins, rem-coin, memo); change >= 0 {
			result = min(result, change+1)
		}
	}

	if result > rem {
		result = -1
	}
	memo[rem-1] = result
	return result
}
