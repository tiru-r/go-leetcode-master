package self_crossing_335

// IsSelfCrossing detects if the path formed by distance array creates self-crossing lines
// Time: O(n), Space: O(1)
func IsSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}
	
	for i := 3; i < n; i++ {
		// Case 1: Fourth line crosses first line (most common case)
		// Pattern: i-3, i-2, i-1, i forms a crossing
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		
		// Case 2: Fifth line crosses first line
		if i >= 4 {
			if distance[i-1] == distance[i-3] && 
			   distance[i] + distance[i-4] >= distance[i-2] {
				return true
			}
		}
		
		// Case 3: Sixth line crosses first line
		if i >= 5 {
			if distance[i-2] >= distance[i-4] && distance[i-3] >= distance[i-1] &&
			   distance[i-1] + distance[i-5] >= distance[i-3] &&
			   distance[i] + distance[i-4] >= distance[i-2] {
				return true
			}
		}
	}
	
	return false
}
