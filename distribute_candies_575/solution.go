package distribute_candies_575

func distributeCandies(candies []int) int {
	half := len(candies) >> 1
	seen := make(map[int]struct{}, half)
	
	for _, candy := range candies {
		seen[candy] = struct{}{}
		if len(seen) == half {
			return half
		}
	}
	
	return len(seen)
}
