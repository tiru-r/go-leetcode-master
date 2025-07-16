package three_sum_15

import "slices"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	slices.Sort(nums)
	var result [][]int

	for i := range len(nums) - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			break
		}

		target := -nums[i]
		left, right := i+1, len(nums)-1

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
