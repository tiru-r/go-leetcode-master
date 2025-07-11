package two_sum_1

// twoSum does not assume that nums is sorted.
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return make([]int, 0, 2)
	}

	numToIndex := make(map[int]int)
	for idx, num := range nums {
		if mIdx, ok := numToIndex[target-num]; ok {
			return []int{mIdx, idx}
		}
		numToIndex[num] = idx
	}

	return make([]int, 0, 2)
}

// twoSumSortedInput assumes that nums is sorted.
func twoSumSortedInput(nums []int, target int) []int {
	if len(nums) < 2 {
		return make([]int, 0, 2)
	}

	front, rear := 0, len(nums)-1
	for front < rear {
		sum := nums[front] + nums[rear]
		switch {
		case sum == target:
			return []int{front, rear}
		case sum < target:
			front++
		default:
			rear--
		}
	}

	return make([]int, 0, 2)
}
