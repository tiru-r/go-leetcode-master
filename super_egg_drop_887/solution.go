package super_egg_drop_887

// SuperEggDrop finds minimum moves to determine critical floor with certainty
// Uses optimized DP with space complexity O(k) and time complexity O(k*moves)
// Time: O(k*sqrt(n)), Space: O(k)
func SuperEggDrop(k int, n int) int {
	if n <= 1 || k == 1 {
		return n
	}
	
	// dp[i] = max floors we can check with i eggs and current number of moves
	dp := make([]int, k+1)
	
	moves := 0
	for dp[k] < n {
		moves++
		// Process backwards to avoid overwriting needed values
		for eggs := k; eggs > 0; eggs-- {
			// If egg breaks: dp[eggs-1] floors with eggs-1 eggs
			// If egg doesn't break: dp[eggs] floors with same eggs
			// Total floors we can check: dp[eggs-1] + 1 + dp[eggs]
			dp[eggs] += dp[eggs-1] + 1
		}
	}
	
	return moves
}

// SuperEggDropMemo alternative implementation using memoization
// More intuitive but slightly less efficient for large inputs
func SuperEggDropMemo(k int, n int) int {
	memo := make(map[[2]int]int)
	return eggDropHelper(k, n, memo)
}

func eggDropHelper(k, n int, memo map[[2]int]int) int {
	if n <= 1 || k == 1 {
		return n
	}
	
	key := [2]int{k, n}
	if val, exists := memo[key]; exists {
		return val
	}
	
	result := n // worst case: linear search
	
	// Binary search for optimal floor
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2
		
		// If egg breaks at floor mid
		broken := eggDropHelper(k-1, mid-1, memo)
		// If egg doesn't break at floor mid  
		notBroken := eggDropHelper(k, n-mid, memo)
		
		// Take worst case + 1 move for current drop
		worstCase := max(broken, notBroken) + 1
		result = min(result, worstCase)
		
		// Adjust binary search bounds
		if broken > notBroken {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	
	memo[key] = result
	return result
}
