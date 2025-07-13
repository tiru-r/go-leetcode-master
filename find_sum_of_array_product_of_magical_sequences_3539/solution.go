package find_sum_of_array_product_of_magical_sequences_3539

const MOD = 1_000_000_007

// FindSumOfArrayProductOfMagicalSequences calculates sum of products for all magical sequences
// Time: O(n * k), Space: O(k) where k is the maximum value in nums
func FindSumOfArrayProductOfMagicalSequences(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k <= 0 {
		return 0
	}
	
	// Find maximum value to determine DP array size
	maxVal := 0
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	
	// dp[i][j] = number of magical sequences of length i ending with value j
	prev := make([]int, maxVal+1)
	curr := make([]int, maxVal+1)
	
	// Initialize: sequences of length 1
	for _, num := range nums {
		prev[num] = (prev[num] + 1) % MOD
	}
	
	if k == 1 {
		sum := 0
		for _, num := range nums {
			sum = (sum + num) % MOD
		}
		return sum
	}
	
	// Build sequences of length 2 to k
	for length := 2; length <= k; length++ {
		// Clear current DP array
		for i := range curr {
			curr[i] = 0
		}
		
		// For each number in nums, try to extend existing sequences
		for _, num := range nums {
			for prevVal := 1; prevVal <= maxVal; prevVal++ {
				if prev[prevVal] > 0 && isValidExtension(prevVal, num) {
					curr[num] = (curr[num] + prev[prevVal]) % MOD
				}
			}
		}
		
		// Swap arrays
		prev, curr = curr, prev
	}
	
	// Calculate sum of products for all valid magical sequences
	result := 0
	productSum := make([]int, maxVal+1)
	
	// Initialize product sums for length 1
	for _, num := range nums {
		productSum[num] = (productSum[num] + num) % MOD
	}
	
	if k == 1 {
		for _, prod := range productSum {
			result = (result + prod) % MOD
		}
		return result
	}
	
	// Recalculate with product tracking
	return calculateProductSum(nums, k, maxVal)
}

// isValidExtension checks if we can extend a sequence ending with prevVal by adding num
func isValidExtension(prevVal, num int) bool {
	// A magical sequence is one where each element is strictly greater than the previous
	return num > prevVal
}

// calculateProductSum calculates the sum of products for all magical sequences
func calculateProductSum(nums []int, k int, maxVal int) int {
	// dp[j] = sum of products of all magical sequences ending with value j
	prev := make([]int, maxVal+1)
	curr := make([]int, maxVal+1)
	
	// Initialize: sequences of length 1
	for _, num := range nums {
		prev[num] = (prev[num] + num) % MOD
	}
	
	if k == 1 {
		sum := 0
		for _, prod := range prev {
			sum = (sum + prod) % MOD
		}
		return sum
	}
	
	// Build sequences of length 2 to k
	for length := 2; length <= k; length++ {
		// Clear current DP array
		for i := range curr {
			curr[i] = 0
		}
		
		// For each number in nums, try to extend existing sequences
		for _, num := range nums {
			for prevVal := 1; prevVal < num && prevVal <= maxVal; prevVal++ {
				if prev[prevVal] > 0 {
					// Add num * (product of previous sequence)
					contribution := (prev[prevVal] * num) % MOD
					curr[num] = (curr[num] + contribution) % MOD
				}
			}
		}
		
		// Swap arrays
		prev, curr = curr, prev
	}
	
	// Sum all products
	result := 0
	for _, prod := range prev {
		result = (result + prod) % MOD
	}
	
	return result
}
