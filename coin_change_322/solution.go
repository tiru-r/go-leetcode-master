package coin_change_322

import (
	"slices" // slices package for sorting, available in Go 1.21+
)

// coinChange calculates the minimum number of coins needed to make a given amount
// using the provided coin denominations. It uses an optimized bottom-up dynamic
// programming approach.
//
// Parameters:
//
//	coins []int: A slice of positive integers representing available coin denominations.
//	             Assumed to be non-empty and contain positive values.
//	amount int: The target amount of money to make change for.
//	            Assumed to be a non-negative integer.
//
// Returns:
//
//	int: The minimum number of coins required to make the amount.
//	     Returns -1 if the amount cannot be made with the given coins.
//
// Time Complexity: O(amount * len(coins)) for the DP loop, plus O(len(coins) * log(len(coins)))
//
//	for sorting coins.
//
// Space Complexity: O(amount) for the DP array.
func coinChange(coins []int, amount int) int {
	// Handle base case: no coins needed for zero amount
	if amount == 0 {
		return 0
	}

	// Handle edge case: no coins provided or amount is negative (though problem implies non-negative amount)
	if len(coins) == 0 || amount < 0 {
		return -1
	}

	// Create a copy of coins to avoid modifying the input slice.
	// This is good practice to maintain function purity.
	coinsCopy := make([]int, len(coins))
	copy(coinsCopy, coins)

	// Sort coins in ascending order. While descending also works,
	// ascending is often more conventional for DP problems building up from smaller values.
	slices.Sort(coinsCopy)

	// Prune coins larger than the target amount. This is an effective optimization
	// to reduce iterations in the inner DP loop.
	validCoins := coinsCopy[:0] // Re-slice to create a zero-length slice backed by coinsCopy
	for _, coin := range coinsCopy {
		if coin <= amount {
			validCoins = append(validCoins, coin)
		}
	}

	// If after pruning, no valid coins remain (e.g., all coins are larger than the amount),
	// then the amount cannot be made.
	if len(validCoins) == 0 {
		return -1
	}
	coins = validCoins // Use the pruned and sorted coins for DP calculation

	// Initialize DP array with a sentinel value representing an unreachable state.
	// `maxInt` is used to avoid potential overflow issues if a very large number
	// were added to it, and ensures it's distinguishable from valid counts.
	// dp[i] will store the minimum coins needed for amount i.
	const maxInt = int(^uint(0) >> 1) // Represents a very large integer (effectively infinity)
	dp := make([]int, amount+1)
	// Initialize all dp states (except dp[0]) to maxInt. dp[0] is already 0 (by default for int slice).
	for i := 1; i <= amount; i++ {
		dp[i] = maxInt
	}

	// Fill the DP table using a bottom-up approach.
	// 'a' represents the current amount being considered.
	for a := 1; a <= amount; a++ {
		// Iterate over each valid coin denomination.
		for _, coin := range coins {
			// If the current coin can be used to form amount 'a' (i.e., coin <= a)
			// AND the subproblem (a - coin) was reachable (dp[a-coin] is not maxInt),
			// then we can potentially update dp[a].
			if coin <= a && dp[a-coin] != maxInt {
				// Update dp[a] with the minimum of its current value and
				// 1 + dp[a-coin] (which represents using the current coin plus
				// the minimum coins needed for the remaining amount).
				// Go 1.21+ provides a built-in min function.
				dp[a] = min(dp[a], dp[a-coin]+1)
			}
		}
		// Early termination optimization: If the target amount 'amount' has been reached
		// and a valid solution has been found for it, we can break early as no further
		// iterations will change dp[amount].
		if a == amount && dp[a] != maxInt {
			break
		}
	}

	// After filling the DP table, check if the target amount was reachable.
	// If dp[amount] is still maxInt, it means the amount cannot be made with the given coins.
	if dp[amount] == maxInt {
		return -1
	}
	// Otherwise, return the minimum number of coins found for the target amount.
	return dp[amount]
}
