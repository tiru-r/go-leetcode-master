package majority_element_ii_229

// majorityElement returns every value that appears more than ⌊n/3⌋ times.
// Extended Boyer-Moore majority vote: two passes, O(n) time, O(1) space.
func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	// Phase 1: find up to two candidates using Boyer-Moore
	var c1, c2, cnt1, cnt2 int
	for _, v := range nums {
		if cnt1 > 0 && v == c1 {
			cnt1++
		} else if cnt2 > 0 && v == c2 {
			cnt2++
		} else if cnt1 == 0 {
			c1, cnt1 = v, 1
		} else if cnt2 == 0 {
			c2, cnt2 = v, 1
		} else {
			cnt1--
			cnt2--
		}
	}

	// Phase 2: verify candidates and handle duplicates
	threshold := len(nums) / 3
	ans := make([]int, 0, 2)
	
	// Prepare candidates for verification, avoiding duplicates
	candidates := make([]int, 0, 2)
	if cnt1 > 0 {
		candidates = append(candidates, c1)
	}
	if cnt2 > 0 && c2 != c1 {
		candidates = append(candidates, c2)
	}
	
	// Count each candidate in a single pass
	for _, candidate := range candidates {
		cnt := 0
		for _, v := range nums {
			if v == candidate {
				cnt++
			}
		}
		if cnt > threshold {
			ans = append(ans, candidate)
		}
	}
	return ans
}
