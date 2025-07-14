package move_zeroes_283

func moveZeroes(nums []int) {
	left := 0
	for right := range nums {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
