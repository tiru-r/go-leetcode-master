package move_zeroes_283

// Ex: [1,0,0,3,12]
//
//	    i
//	p
//
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

