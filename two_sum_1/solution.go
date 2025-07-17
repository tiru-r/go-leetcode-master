package two_sum_1

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))

	for currIdx, num := range nums {
		if prevIdx, ok := index[target-num]; ok {
			return []int{prevIdx, currIdx}
		}
		index[num] = currIdx
	}
	return nil
}

func twoSumSorted(nums []int, target int) (int, int) {
	i, j := 0, len(nums)-1
	for i < j {
		sum := nums[i] + nums[j]
		switch {
		case sum == target:
			return i, j
		case sum < target:
			i++
		default:
			j--
		}
	}
	return -1, -1
}
