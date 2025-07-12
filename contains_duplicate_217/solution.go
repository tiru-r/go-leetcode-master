package contains_duplicate_217

// ContainsDuplicate checks if any value appears at least twice in the given integer slice.
// It uses a hash map to track encountered numbers, returning true if a duplicate is found.
//
// Time Complexity: O(n), where n is the length of the input slice.
//   - Iterates through the slice once, with O(1) average-case map operations.
//
// Space Complexity: O(n), where n is the length of the input slice.
//   - Stores up to n elements in the map in the worst case (all unique elements).
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums)) // Pre-allocate map with hint for capacity.

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
