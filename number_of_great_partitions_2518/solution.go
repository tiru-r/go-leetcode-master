package number_of_great_partitions_2518

const mod = 1_000_000_007

// WaysToPartition returns the number of ways to pick an index i
// and change nums[i] to k such that the array can be partitioned
// into two non-empty contiguous sub-arrays with equal sums.
func WaysToPartition(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	// Calculate total sum of the array
	total := 0
	for _, value := range nums {
		total += value
	}

	// Count base partitions (without any changes)
	result := 0
	prefix := 0
	for i := 0; i < n-1; i++ {
		prefix += nums[i]
		if prefix*2 == total {
			result++
		}
	}

	// Initialize sliding hash maps for difference tracking
	left := make(map[int]int, n)
	right := make(map[int]int, n)

	// Pre-compute all right-side differences
	prefix = 0
	for i := 0; i < n-1; i++ {
		prefix += nums[i]
		diff := prefix*2 - total
		right[diff]++
	}

	// Sweep through each possible position to change
	prefix = 0
	for i := 0; i < n; i++ {
		delta := k - nums[i]
		// Cache the doubled delta to avoid repeated calculation
		doubleDelta := 2 * delta

		// Mathematical derivation:
		// For equal partition after change: (prefix ± delta)*2 == total
		// This translates to difference requirements:
		//   Left side needs: diff = -2*delta
		//   Right side needs: diff = 2*delta
		result = (result + left[-doubleDelta] + right[doubleDelta]) % mod

		// Move to next position and update sliding window
		if i < n-1 {
			prefix += nums[i]
			diff := prefix*2 - total

			// Transfer difference count from right to left map
			right[diff]--
			// Optimization: keep zero entries to avoid map resize overhead
			left[diff]++
		}
	}

	return result
}
