package search_in_rotated_sorted_array_33

// search finds the index of target in a rotated sorted array in O(log n) time.
// Returns -1 if the target is not found.
func search(nums []int, target int) int {
	// Handle empty array
	if len(nums) == 0 {
		return -1
	}

	// Micro-optimization: Handle single-element array
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2 // Prevents integer overflow

		if nums[mid] == target {
			return mid
		}

		// Check if left half (start to mid) is sorted
		if nums[start] <= nums[mid] {
			// Check if target is in the sorted left half
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { // Right half (mid to end) is sorted
			// Check if target is in the sorted right half
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1 // Target not found
}
