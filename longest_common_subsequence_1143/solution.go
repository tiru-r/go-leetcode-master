package longest_common_subsequence_1143

// Note: study again.
// DP bottoms-up solution. Builds on same idea of recursive solution but
// starts from adding to subsequences from front-to-back. Starting at front
// allows to use previous solutions to build the next solution.
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text2)+1)
	for i := range dp {
		dp[i] = make([]int, len(text1)+1)
	}

	for m := 1; m < len(dp); m++ {
		for n := 1; n < len(dp[m]); n++ {
			if text1[n-1] == text2[m-1] {
				dp[m][n] = 1 + dp[m-1][n-1]
			} else { // text1[n - 1] != text2[m - 1]
				dp[m][n] = max(dp[m][n-1], dp[m-1][n])
			}
		}
	}

	return dp[len(text2)][len(text1)]
}

