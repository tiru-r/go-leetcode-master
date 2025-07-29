package three_sum_15

import "slices"

// threeSum finds all unique triplets that sum to zero using two-pointer technique.
// Time: O(n²), Space: O(1) excluding output.
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	// Sort to enable two-pointer technique and duplicate handling
	slices.Sort(nums)
	
	// Pre-allocate with estimated capacity (typical case: few triplets)
	result := make([][]int, 0, len(nums)/10)

	// Fix first element and use two pointers for remaining pair
	n := len(nums)
	for i := range n - 2 {
		// Skip duplicate first elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early termination: if smallest remaining number > 0, no zero sum possible
		if nums[i] > 0 {
			break
		}

		// Two-pointer search for pair that sums to -nums[i]
		target := -nums[i]
		left, right := i+1, n-1

		// Two-pointer search in sorted subarray
		for left < right {
			currentSum := nums[left] + nums[right]
			
			switch {
			case currentSum == target:
				// Found valid triplet
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Skip duplicate left values
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicate right values
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move both pointers to continue search
				left++
				right--
				
			case currentSum < target:
				left++ // Need larger sum
			default:
				right-- // Need smaller sum
			}
		}
	}

	return result
}
