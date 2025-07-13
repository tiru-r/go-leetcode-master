package find_the_difference_of_two_arrays_2215


// findDifference returns elements unique to each array
// Optimized: O(n+m) time, O(n+m) space using built-in set operations
func findDifference(nums1, nums2 []int) [][]int {
	set1, set2 := buildSets(nums1, nums2)
	
	// Use efficient set operations with pre-allocated slices
	diff1 := make([]int, 0, len(set1))
	diff2 := make([]int, 0, len(set2))
	
	// Single pass difference calculation
	for val := range set1 {
		if !set2[val] {
			diff1 = append(diff1, val)
		}
	}
	
	for val := range set2 {
		if !set1[val] {
			diff2 = append(diff2, val)
		}
	}
	
	return [][]int{diff1, diff2}
}

// buildSets creates sets from both arrays in parallel processing style
func buildSets(nums1, nums2 []int) (map[int]bool, map[int]bool) {
	set1 := make(map[int]bool, len(nums1))
	set2 := make(map[int]bool, len(nums2))
	
	// Use modern range-based approach
	for _, val := range nums1 {
		set1[val] = true
	}
	
	for _, val := range nums2 {
		set2[val] = true
	}
	
	return set1, set2
}
