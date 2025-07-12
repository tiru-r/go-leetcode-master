package coin_change_322

// coinChange calculates the minimum number of coins needed to make up a given amount.
// If the amount cannot be formed by any combination of the coins, it returns -1.
//
// This function uses a dynamic programming (DP) approach:
// dp[i] represents the minimum number of coins required to make up amount 'i'.
//
// Time Complexity: O(amount * len(coins))
//   - The outer loop iterates 'amount' times.
//   - The inner loop iterates 'len(coins)' times.
//
// Space Complexity: O(amount)
//   - A DP array of size 'amount + 1' is used.
func coinChange(coins []int, amount int) int {
	// Base case: If the amount is 0, no coins are needed.
	if amount == 0 {
		return 0
	}

	// Edge cases: If there are no coins or the amount is negative, return -1.
	if len(coins) == 0 || amount < 0 {
		return -1
	}

	// Early exit for single coin case
	if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}

	// Initialize the DP array with amount+1 as the sentinel value
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // Represents an unreachable state
	}

	// Iterate through each possible amount from 1 up to the target amount
	for a := 1; a <= amount; a++ {
		// For each amount, try all coin denominations
		for _, c := range coins {
			if c <= a && dp[a-c] != amount+1 {
				dp[a] = min(dp[a], dp[a-c]+1)
			}
		}
	}

	// If dp[amount] is still amount+1, the amount is unreachable
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
