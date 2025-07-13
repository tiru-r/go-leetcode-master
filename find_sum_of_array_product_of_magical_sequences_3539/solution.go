package find_sum_of_array_product_of_magical_sequences_3539

import (
	"slices"
)

const MOD = 1_000_000_007

// FindSumOfArrayProductOfMagicalSequences calculates sum of products for all magical sequences
// Optimized: Time O(n²*k), Space O(n) using coordinate compression and prefix sums
func FindSumOfArrayProductOfMagicalSequences(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k <= 0 {
		return 0
	}
	
	// Early termination for impossible cases
	if k > n {
		return 0
	}
	
	// Coordinate compression: map values to indices for sparse representation
	unique := make([]int, 0, n)
	seen := make(map[int]bool)
	for _, num := range nums {
		if !seen[num] {
			unique = append(unique, num)
			seen[num] = true
		}
	}
	slices.Sort(unique)
	
	// Map value to compressed index
	compress := make(map[int]int, len(unique))
	for i, val := range unique {
		compress[val] = i
	}
	
	m := len(unique)
	
	// Fast path for k=1
	if k == 1 {
		sum := 0
		for _, num := range nums {
			sum = (sum + num) % MOD
		}
		return sum
	}
	
	// dp[i] = sum of products of magical sequences ending at unique[i]
	prev := make([]int, m)
	curr := make([]int, m)
	
	// Initialize with single elements
	for _, num := range nums {
		idx := compress[num]
		prev[idx] = (prev[idx] + num) % MOD
	}
	
	// Build sequences of length 2 to k using prefix sums for optimization
	for length := 2; length <= k; length++ {
		clear(curr)
		
		// Prefix sum array for fast range queries
		prefixSum := make([]int, m+1)
		for i := range m {
			prefixSum[i+1] = (prefixSum[i] + prev[i]) % MOD
		}
		
		// For each unique value, extend all valid previous sequences
		for i, val := range unique {
			// Find all smaller values using binary search
			pos := lowerBound(unique, val)
			if pos > 0 {
				// Sum of all sequences ending with values < val
				contribution := (prefixSum[pos] * val) % MOD
				curr[i] = (curr[i] + contribution) % MOD
			}
		}
		
		prev, curr = curr, prev
	}
	
	// Sum all final products
	result := 0
	for _, prod := range prev {
		result = (result + prod) % MOD
	}
	
	return result
}

// lowerBound returns the first index where unique[i] >= target
func lowerBound(unique []int, target int) int {
	left, right := 0, len(unique)
	for left < right {
		mid := (left + right) / 2
		if unique[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
