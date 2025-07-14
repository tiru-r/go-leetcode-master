package longest_consecutive_sequence_128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Use empty struct for memory efficiency
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	maxLength := 0
	for num := range set {
		// Only start counting if this is the beginning of a sequence
		if _, exists := set[num-1]; !exists {
			currentNum := num
			currentLength := 1
			
			// Count consecutive numbers
			for {
				if _, exists := set[currentNum+1]; !exists {
					break
				}
				currentNum++
				currentLength++
			}
			
			maxLength = max(maxLength, currentLength)
		}
	}
	
	return maxLength
}
