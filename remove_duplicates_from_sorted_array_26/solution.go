package remove_duplicates_from_sorted_array_26

// Modern solution using two-pointer technique
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIndex := 1
	for readIndex := 1; readIndex < len(nums); readIndex++ {
		if nums[readIndex] != nums[readIndex-1] {
			nums[writeIndex] = nums[readIndex]
			writeIndex++
		}
	}

	return writeIndex
}
