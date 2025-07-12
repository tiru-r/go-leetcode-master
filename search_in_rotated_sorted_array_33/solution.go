package search_in_rotated_sorted_array_33

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// if no rotation happened to the sorted nums, the
	// just binary search over the entire nums
	if nums[0] <= nums[len(nums)-1] {
		// no rotation happened at all
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	// binary search for the inflection point
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := start + ((end - start) / 2)

		// if mid is smallest at inflection
		if mid > 0 && nums[mid] < nums[mid-1] {
			// bring mid back to largest
			mid--
			break
		}

		// if mid is largest at inflection
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			break
		}

		if nums[0] < nums[mid] {
			// inflection point of right of mid
			start = mid
		} else {
			// inflection point of left of mid
			end = mid
		}
	}

	// if the target is less than nums at 0, then it falls on
	// the right-hand side of the inflection point
	if target < nums[0] {
		return binarySearch(nums, target, mid+1, len(nums)-1)
	}

	// if the target is greater than nums at the last index, then
	// it falls on the left-hand side of the inflection point
	if target > nums[len(nums)-1] {
		// target is on left-hand side of inflection point
		return binarySearch(nums, target, 0, mid)
	}

	return -1
}

func binarySearch(nums []int, target, start, end int) int {
	for start <= end {
		mid := start + ((end - start) / 2)
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

