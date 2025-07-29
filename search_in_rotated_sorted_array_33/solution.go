package search_in_rotated_sorted_array_33

// search returns the index of target in a rotated sorted array using modified binary search.
// Time: O(log n), Space: O(1). Returns -1 if target is not present.
// Algorithm: At each step, one half of the array is guaranteed to be sorted.
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	
	for lo <= hi {
		mid := lo + (hi-lo)>>1 // Bit shift for fast division, prevents overflow
		
		if nums[mid] == target {
			return mid
		}

		// Determine which half is sorted by comparing endpoints
		if nums[lo] <= nums[mid] {
			// Left half [lo..mid] is sorted
			if target >= nums[lo] && target < nums[mid] {
				hi = mid - 1 // Target is in sorted left half
			} else {
				lo = mid + 1 // Target must be in right half
			}
		} else {
			// Right half [mid..hi] is sorted (rotation point is in left half)
			if target > nums[mid] && target <= nums[hi] {
				lo = mid + 1 // Target is in sorted right half
			} else {
				hi = mid - 1 // Target must be in left half
			}
		}
	}
	return -1 // Target not found
}
