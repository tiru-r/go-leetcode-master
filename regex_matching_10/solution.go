package regex_matching_10

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([]bool, n+1)
	prev := make([]bool, n+1)

	dp[0] = true
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[j] = dp[j-2]
		}
	}

	for i := 1; i <= m; i++ {
		prev, dp = dp, prev
		dp[0] = false

		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[j] = dp[j-2] || (matches(s[i-1], p[j-2]) && prev[j])
			} else {
				dp[j] = matches(s[i-1], p[j-1]) && prev[j-1]
			}
		}
	}

	return dp[n]
}

func matches(sc, pc byte) bool {
	return pc == '.' || sc == pc
}
