package remove_duplicates_from_sorted_array_26

import "slices"

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

// Alternative using slices.Compact (Go 1.21+)
func removeDuplicatesCompact(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	
	// slices.Compact removes consecutive duplicate elements in-place
	compacted := slices.Compact(nums)
	return len(compacted)
}

// Original implementation (kept for comparison)
func removeDuplicatesOriginal(nums []int) int {
	i, j, u := 0, 0, 1
	for j < len(nums)-1 {
		j = i
		for j < len(nums)-1 && nums[i] == nums[j] {
			j++
		}
		if nums[i] != nums[j] {
			u++
		}

		for k := j - 1; k > i; k-- {
			nums[k] = nums[j]
		}

		i++
	}

	return u
}
