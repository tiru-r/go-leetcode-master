package merge_sorted_array_88

// O(m+n) solution
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	n1, n2, p := m-1, n-1, m+n-1
	for n2 > -1 {
		if n1 < 0 || nums2[n2] >= nums1[n1] {
			nums1[p] = nums2[n2]
			n2--
		} else {
			nums1[p] = nums1[n1]
			n1--
		}
		p--
	}
}
