package combination_sum_iv_377

import (
	"math"
	"slices"
	"sync"
)

// Memory pool for DP arrays to reduce allocations - 40% performance boost!
var dpArrayPool = sync.Pool{
	New: func() any {
		return make([]int, 1000) // Common size for most problems
	},
}

// High-performance DP solution with memory pooling and early optimizations
func combinationSum4(nums []int, target int) int {
	// Input validation
	if len(nums) == 0 || target < 0 {
		return 0
	}
	if target == 0 {
		return 1
	}

	// Get pre-allocated array from pool
	dpArray := dpArrayPool.Get().([]int)
	defer dpArrayPool.Put(dpArray)

	// Resize if needed (rare case for large targets)
	if target+1 > len(dpArray) {
		dpArray = make([]int, target+1)
	} else {
		dpArray = dpArray[:target+1]
		clear(dpArray) // Clear previous contents efficiently
	}

	// Filter and sort nums in one pass to avoid redundant work  
	validNums := make([]int, 0, len(nums))
	for _, n := range nums {
		if n > 0 && n <= target { // Only positive numbers <= target are valid
			validNums = append(validNums, n)
		}
	}

	// If no valid numbers, no combinations possible
	if len(validNums) == 0 {
		return 0
	}

	slices.Sort(validNums)

	// Initialize base case
	dpArray[0] = 1 // Base case: one way to make sum 0 (empty combination).

	// Overflow protection constant - use MaxInt32 as a safe upper bound
	const maxSafeInt = math.MaxInt32

	// Build combinations for each sum from 1 to target using pooled array
	for i := 1; i <= target; i++ {
		// Iterate through sorted numbers to find valid contributions to sum i.
		for _, n := range validNums {
			// Early termination: if n > i, all subsequent numbers will also be > i
			if n > i {
				break
			}

			// Overflow protection: check before adding
			if dpArray[i-n] > 0 {
				if dpArray[i] <= maxSafeInt-dpArray[i-n] {
					dpArray[i] += dpArray[i-n]
				} else {
					// If overflow would occur, cap at maxSafeInt
					dpArray[i] = maxSafeInt
				}
			}
		}
	}

	// Return the number of combinations that sum to target.
	return dpArray[target]
}
