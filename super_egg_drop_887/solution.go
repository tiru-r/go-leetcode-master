package super_egg_drop_887

// superEggDrop returns the minimum number of moves needed
// to determine the critical floor f with certainty.
func superEggDrop(k int, n int) int {
	// dp[i] = max floors we can cover with i eggs and the current number of moves
	dp := make([]int, k+1)

	moves := 0
	for dp[k] < n {
		moves++
		// iterate backwards so we reuse the previous state
		for eggs := k; eggs > 0; eggs-- {
			dp[eggs] += dp[eggs-1] + 1
		}
	}
	return moves
}
