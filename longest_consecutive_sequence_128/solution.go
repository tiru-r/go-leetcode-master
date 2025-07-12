package longest_consecutive_sequence_128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Build a set for O(1) existence checks
	set := make(map[int]bool, len(nums))
	for _, v := range nums {
		set[v] = true
	}

	longest := 0
	for v := range set {
		// only start counting if v-1 is NOT in the set
		if !set[v-1] {
			length := 1
			for cur := v + 1; set[cur]; cur++ {
				length++
			}
			if length > longest {
				longest = length
			}
		}
	}
	return longest
}
