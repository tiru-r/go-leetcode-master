package find_minimum_in_rotated_sorted_array_153

func findMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1) // avoid overflow
		if nums[mid] > nums[hi] {    // pivot on right
			lo = mid + 1
		} else { // pivot on left (or mid == pivot)
			hi = mid
		}
	}
	return nums[lo]
}
