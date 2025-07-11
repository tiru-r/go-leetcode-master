package contains_duplicate_217

import "slices"

// Original optimized solution using map for O(n) time complexity
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, n := range nums {
		if _, exists := m[n]; exists {
			return true
		}

		m[n] = struct{}{}
	}

	return false
}

// Alternative implementation using slices.Sort for O(n log n) solution
func containsDuplicateSort(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	// Create a copy to avoid modifying the original
	sorted := slices.Clone(nums)
	slices.Sort(sorted)

	// Check adjacent elements for duplicates
	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			return true
		}
	}

	return false
}

/*
Note: bit shifting technique only worked for n > -1 in nums
func containsDuplicate(nums []int) bool {
	var dups uint64
	for _, n := range nums {
		exists := dups & (1 << uint64(n))
		if exists == 0 {
			dups |= 1 << uint64(n)
		} else {
			return true
		}
	}

	return false
}
*/
