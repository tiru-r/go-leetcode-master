package three_sum_15

import "slices"

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}

	slices.Sort(nums)
	var result [][]int

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early termination: if smallest element is positive, no solution exists
		if nums[i] > 0 {
			break
		}

		target := -nums[i]
		left, right := i+1, n-1

		for left < right {
			sum := nums[left] + nums[right]
			switch {
			case sum == target:
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			case sum < target:
				left++
			default:
				right--
			}
		}
	}

	return result
}
