package house_robber_198

// rob returns maximum money robbed using optimized space-efficient DP
// Optimized: O(n) time, O(1) space with Go 1.24 modern syntax
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// DP state: rob = max money if we rob current, skip = max if we skip current
	rob, skip := 0, 0
	
	for _, money := range nums {
		// Choice: rob current + previous skip, or keep previous rob
		rob, skip = max(skip+money, rob), rob
	}
	
	return rob
}
