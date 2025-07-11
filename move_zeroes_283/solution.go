package move_zeroes_283

// Ex: [1,0,0,3,12]
//
//	    i
//	p
// Modern solution using two-pointer technique with clean swapping
func moveZeroes(nums []int) {
	writeIndex := 0
	for readIndex := 0; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != 0 {
			nums[writeIndex], nums[readIndex] = nums[readIndex], nums[writeIndex]
			writeIndex++
		}
	}
}

// Alternative implementation preserving order efficiently
func moveZeroes2(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			if left != right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
	}
}

// Original implementation (kept for comparison)
func moveZeroesOriginal(nums []int) {
	placement := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			hold := nums[i]
			nums[i] = nums[placement]
			nums[placement] = hold
			placement++
		}
	}
}
