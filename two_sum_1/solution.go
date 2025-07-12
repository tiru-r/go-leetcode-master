package two_sum_1

// TwoSum returns the indices of the two numbers in nums that add up to target.
// Exactly one solution is assumed to exist; if the caller violates this
// guarantee, an empty slice is returned.
//
// Time complexity: O(n) average, O(n²) worst (pathological hash collisions)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	// Pre-size the map to avoid rehashing.
	index := make(map[int]int, len(nums))

	for currIdx, num := range nums {
		if prevIdx, ok := index[target-num]; ok {
			return []int{prevIdx, currIdx}
		}
		index[num] = currIdx
	}
	return []int{}
}

// TwoSumSorted assumes nums is in non-decreasing order.
// Time: O(n), Space: O(1)
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
	return -1, -1 // sentinel: no solution
}
