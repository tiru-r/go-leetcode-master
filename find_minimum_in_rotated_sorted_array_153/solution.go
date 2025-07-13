package find_minimum_in_rotated_sorted_array_153

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	
	return nums[left]
}
