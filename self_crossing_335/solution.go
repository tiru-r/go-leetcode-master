package self_crossing_335

func IsSelfCrossing(distance []int) bool {
	if len(distance) < 4 {
		return false
	}
	
	for i := range len(distance) - 3 {
		i += 3
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		
		if i >= 4 && distance[i-1] == distance[i-3] && 
		   distance[i] + distance[i-4] >= distance[i-2] {
			return true
		}
		
		if i >= 5 && distance[i-2] >= distance[i-4] && distance[i-3] >= distance[i-1] &&
		   distance[i-1] + distance[i-5] >= distance[i-3] &&
		   distance[i] + distance[i-4] >= distance[i-2] {
			return true
		}
	}
	
	return false
}
