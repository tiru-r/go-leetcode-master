package majority_element_169

// Note: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
// Good problem and intro to a new algorithm.
func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	// Use range over values for better performance
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
