package coin_change_322

import (
	"math"
)

// coinChange returns the minimum number of coins needed to make `amount`.
// Returns -1 if the amount cannot be formed.
// Time  : O(amount * len(coins))
// Space : O(amount)
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 || amount < 0 {
		return -1
	}

	// Filter out coins larger than the target (optional micro-optimisation)
	validCoins := make([]int, 0, len(coins))
	for _, c := range coins {
		if c <= amount {
			validCoins = append(validCoins, c)
		}
	}
	if len(validCoins) == 0 {
		return -1
	}

	// DP array: dp[i] = minimum coins for amount i
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for a := 1; a <= amount; a++ {
		for _, c := range validCoins {
			if c <= a && dp[a-c] != math.MaxInt {
				dp[a] = min(dp[a], dp[a-c]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
