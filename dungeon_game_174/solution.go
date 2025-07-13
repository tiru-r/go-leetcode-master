package dungeon_game_174

// CalculateMinimumHP finds the minimum initial health needed to rescue the princess
// Time: O(m*n), Space: O(n) - optimized space using 1D array
func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	
	m, n := len(dungeon), len(dungeon[0])
	
	// Use 1D array for space optimization - represents one row
	dp := make([]int, n+1)
	
	// Initialize boundaries - need infinite health to go out of bounds
	for j := range n + 1 {
		dp[j] = 1<<30
	}
	dp[n] = 1 // Right boundary
	
	// Fill DP table from bottom-right to top-left
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				// At princess location - need enough to survive this cell
				dp[j] = max(1, 1-dungeon[i][j])
			} else {
				// Minimum health needed from next cells
				need := min(dp[j], dp[j+1]) - dungeon[i][j]
				dp[j] = max(1, need)
			}
		}
	}
	
	return dp[0]
}
