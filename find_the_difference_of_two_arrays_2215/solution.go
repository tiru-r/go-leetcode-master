package find_the_difference_of_two_arrays_2215

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := sliceToSet(nums1)
	set2 := sliceToSet(nums2)
	return [][]int{
		setDifference(set1, set2),
		setDifference(set2, set1),
	}
}

// setDifference returns the keys that are in s1 but not in s2.
func setDifference(s1, s2 map[int]struct{}) []int {
	res := make([]int, 0, len(s1))
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			res = append(res, k)
		}
	}
	return res
}

// sliceToSet converts a slice to a set (map[int]struct{}).
func sliceToSet(a []int) map[int]struct{} {
	s := make(map[int]struct{}, len(a))
	for _, v := range a {
		s[v] = struct{}{}
	}
	return s
}
