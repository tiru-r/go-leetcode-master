package find_the_difference_of_two_arrays_2215

import (
	"maps"
	"slices"
)

// Original optimized solution using map-based set operations
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := sliceToSet(nums1)
	set2 := sliceToSet(nums2)
	return [][]int{
		setDifference(set1, set2),
		setDifference(set2, set1),
	}
}

func setDifference(s1, s2 map[int]struct{}) []int {
	var res []int
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			res = append(res, k)
		}
	}

	return res
}

func sliceToSet(a []int) map[int]struct{} {
	s := make(map[int]struct{})
	for _, v := range a {
		s[v] = struct{}{}
	}
	return s
}

// Alternative solution using slices.Contains (less efficient O(n²) but more readable)
func findDifferenceSlices(nums1 []int, nums2 []int) [][]int {
	// Remove duplicates from both arrays
	unique1 := removeDuplicates(nums1)
	unique2 := removeDuplicates(nums2)

	var diff1, diff2 []int

	// Find elements in nums1 but not in nums2
	for _, num := range unique1 {
		if !slices.Contains(unique2, num) {
			diff1 = append(diff1, num)
		}
	}

	// Find elements in nums2 but not in nums1
	for _, num := range unique2 {
		if !slices.Contains(unique1, num) {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}
}

func removeDuplicates(nums []int) []int {
	seen := make(map[int]struct{})
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		if _, exists := seen[num]; !exists {
			result = append(result, num)
			seen[num] = struct{}{}
		}
	}

	return result
}

// Modern version using maps package
func findDifferenceWithMaps(nums1 []int, nums2 []int) [][]int {
	set1 := sliceToSet(nums1)
	set2 := sliceToSet(nums2)

	// Create copies for manipulation
	diff1Set := maps.Clone(set1)
	diff2Set := maps.Clone(set2)

	// Remove common elements
	maps.DeleteFunc(diff1Set, func(k int, _ struct{}) bool {
		_, exists := set2[k]
		return exists
	})

	maps.DeleteFunc(diff2Set, func(k int, _ struct{}) bool {
		_, exists := set1[k]
		return exists
	})

	// Convert map keys to slices using modern maps.Keys() with slices.Collect
	diff1 := slices.Collect(maps.Keys(diff1Set))
	diff2 := slices.Collect(maps.Keys(diff2Set))

	return [][]int{diff1, diff2}
}
