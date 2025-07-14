package intersection_of_two_arrays_ii_350

// intersect finds multiset intersection using optimized frequency counting
// Optimized: O(m+n) time, O(min(m,n)) space with Go 1.24 modern syntax
func intersect(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	// Use smaller array for frequency map to minimize memory usage
	smaller, larger := nums1, nums2
	if len(nums2) < len(nums1) {
		smaller, larger = nums2, nums1
	}

	// Count frequencies in smaller array
	freq := make(map[int]int, len(smaller))
	for _, num := range smaller {
		freq[num]++
	}

	// Collect intersection by scanning larger array
	result := make([]int, 0, len(smaller))
	for _, num := range larger {
		if count := freq[num]; count > 0 {
			result = append(result, num)
			freq[num]--
		}
	}
	
	return result
}
