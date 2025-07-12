package contains_duplicate_217

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

