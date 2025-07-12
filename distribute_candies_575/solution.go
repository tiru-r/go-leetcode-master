package distribute_candies_575

func distributeCandies(candies []int) int {
	// Fast uniqueness with a map[int]struct{}
	unique := make(map[int]struct{}, len(candies))
	for _, v := range candies {
		unique[v] = struct{}{}
	}
	return min(len(unique), len(candies)/2)
}
