package contains_duplicate_217

// ContainsDuplicate reports whether any integer appears at least twice.
//
//	Time:  O(n)
//	Space: O(n)
//	GC:    zero allocations after map growth (Go 1.21+)
func containsDuplicate(nums []int) bool {
	// Pre-size the map to avoid re-hashing; worst-case = all unique.
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
