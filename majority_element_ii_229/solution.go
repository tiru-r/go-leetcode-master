package majority_element_ii_229

// majorityElement finds all elements that appear more than ⌊n/3⌋ times.
// Uses Boyer-Moore majority vote algorithm extended for k=2 candidates.
// Time: O(n), Space: O(1)
func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// Phase 1: Find potential candidates using Boyer-Moore
	candidate1, candidate2 := 0, 1 // Use different initial values to avoid collision
	count1, count2 := 0, 0

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1, count1 = num, 1
		} else if count2 == 0 {
			candidate2, count2 = num, 1
		} else {
			count1--
			count2--
		}
	}

	// Phase 2: Verify candidates actually appear > n/3 times
	count1, count2 = 0, 0
	threshold := len(nums) / 3

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	result := make([]int, 0, 2)
	if count1 > threshold {
		result = append(result, candidate1)
	}
	if count2 > threshold && candidate2 != candidate1 {
		result = append(result, candidate2)
	}

	return result
}
