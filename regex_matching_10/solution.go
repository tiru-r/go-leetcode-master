package regex_matching_10

const (
	dot   = '.' // wildcard
	star  = '*' // Kleene star
	empty = ""  // sentinel for human readability
)

// isMatch reports whether the entire string s matches the pattern p.
// The pattern may contain:
//   - '.'  which matches any single character
//   - '*'  which matches zero or more of the preceding element
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// dp[j] == true  ⇔  s[:i] matches p[:j]
	dp := make([]bool, n+1)
	prev := make([]bool, n+1)

	// Empty string vs empty pattern → match
	dp[0] = true

	// Empty string vs patterns like a*, a*b*, a*b*c* ...
	for j := 2; j <= n; j += 2 {
		if p[j-1] == star {
			dp[j] = dp[j-2]
		}
	}

	for i := 1; i <= m; i++ {
		// Swap roles: prev is dp[i-1], dp becomes dp[i]
		prev, dp = dp, prev
		dp[0] = false // no pattern can match a non-empty string

		for j := 1; j <= n; j++ {
			switch p[j-1] {
			case star:
				// zero occurrences OR one-or-more occurrences
				dp[j] = dp[j-2] || // skip the x*
					(charMatch(s[i-1], p[j-2]) && prev[j]) // use the star
			default:
				dp[j] = charMatch(s[i-1], p[j-1]) && prev[j-1]
			}
		}
	}
	return dp[n]
}

// charMatch returns true if the pattern char pc matches the string char sc.
func charMatch(sc byte, pc byte) bool {
	return pc == dot || sc == pc
}
