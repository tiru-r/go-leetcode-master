package super_egg_drop_887

func SuperEggDrop(k, n int) int {
	if n <= 1 || k == 1 {
		return n
	}
	
	dp := make([]int, k+1)
	
	for moves := 0; dp[k] < n; moves++ {
		for eggs := range k {
			dp[k-eggs] += dp[k-eggs-1] + 1
		}
		if dp[k] >= n {
			return moves + 1
		}
	}
	
	return 0
}
