package number_of_great_partitions_2518

// WaysToPartition counts the number of great partitions after changing at most one element
// Time: O(n), Space: O(n)
func WaysToPartition(nums []int, k int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	
	total := 0
	for _, num := range nums {
		total += num
	}
	
	// Count valid partitions without changing any element
	result := 0
	prefixSum := 0
	for i := 0; i < n-1; i++ {
		prefixSum += nums[i]
		if prefixSum*2 == total {
			result++
		}
	}
	
	// Count differences for each partition point
	leftDiffs := make(map[int]int)  // differences for partitions to the left
	rightDiffs := make(map[int]int) // differences for partitions to the right
	
	// Initialize rightDiffs with all partition differences
	prefixSum = 0
	for i := 0; i < n-1; i++ {
		prefixSum += nums[i]
		diff := prefixSum*2 - total
		rightDiffs[diff]++
	}
	
	// Try changing each element
	prefixSum = 0
	for i := 0; i < n; i++ {
		delta := k - nums[i]
		
		// Check how many partitions become valid when we change nums[i] to k
		// For left partitions: need prefixSum + delta == total - prefixSum - delta
		// => 2*(prefixSum + delta) == total => 2*prefixSum + 2*delta == total
		// => 2*delta == total - 2*prefixSum => delta == (total - 2*prefixSum)/2
		// We have: 2*prefixSum - total + 2*delta == 0 => diff == -2*delta
		result += leftDiffs[-2*delta]
		
		// For right partitions: need prefixSum == total - prefixSum + delta  
		// => 2*prefixSum == total + delta => diff == 2*delta
		result += rightDiffs[2*delta]
		
		// Update for next iteration
		if i < n-1 {
			prefixSum += nums[i]
			diff := prefixSum*2 - total
			
			// Move this partition from right to left
			rightDiffs[diff]--
			if rightDiffs[diff] == 0 {
				delete(rightDiffs, diff)
			}
			leftDiffs[diff]++
		}
	}
	
	return result
}