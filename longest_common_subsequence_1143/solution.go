package longest_common_subsequence_1143

func longestCommonSubsequence(text1, text2 string) int {
	n, m := len(text1), len(text2)
	if n == 0 || m == 0 {
		return 0
	}

	// Ensure text1 is the shorter string for space optimization
	if n > m {
		text1, text2 = text2, text1
		n, m = m, n
	}

	// Space-optimized DP: O(min(n,m)) space instead of O(n*m)
	prev := make([]int, n+1)
	curr := make([]int, n+1)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text2[i-1] == text1[j-1] {
				curr[j] = 1 + prev[j-1]
			} else {
				curr[j] = max(prev[j], curr[j-1])
			}
		}
		prev, curr = curr, prev
	}
	
	return prev[n]
}
