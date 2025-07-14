package majority_element_169

// majorityElement finds the majority element using Boyer-Moore algorithm.
// Time: O(n), Space: O(1)
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
