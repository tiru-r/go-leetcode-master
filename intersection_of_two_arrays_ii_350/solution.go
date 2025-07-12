package intersection_of_two_arrays_ii_350

// intersect returns the multiset intersection of nums1 and nums2.
// Time: O(len(nums1)+len(nums2))   Space: O(min(len(nums1),len(nums2)))
func intersect(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}

	// Count frequencies in the smaller slice for better memory usage.
	var base, scan []int
	if len(nums1) < len(nums2) {
		base, scan = nums1, nums2
	} else {
		base, scan = nums2, nums1
	}

	count := make(map[int]int, len(base))
	for _, v := range base {
		count[v]++
	}

	res := make([]int, 0, min(len(base), len(scan)))
	for _, v := range scan {
		if cnt := count[v]; cnt > 0 {
			res = append(res, v)
			count[v]--
		}
	}
	return res
}
