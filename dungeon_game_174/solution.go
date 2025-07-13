package dungeon_game_174

func CalculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n+1)
	
	for i := range dp {
		dp[i] = 1 << 30
	}
	dp[n-1] = 1
	
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[j] = max(1, 1-dungeon[i][j])
			} else {
				dp[j] = max(1, min(dp[j], dp[j+1])-dungeon[i][j])
			}
		}
	}
	
	return dp[0]
}
