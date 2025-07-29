package find_sum_of_array_product_of_magical_sequences_3539

import "sort"

const MOD = 1_000_000_007

// FindSumOfArrayProductOfMagicalSequences returns the required sum modulo MOD.
// Time: O(n²·k)   Space: O(n)
func FindSumOfArrayProductOfMagicalSequences(nums []int, k int) int {
	n := len(nums)
	if n == 0 || k <= 0 || k > n {
		return 0
	}

	// collect and sort unique values
	seen := make(map[int]struct{}, n)
	unique := make([]int, 0, n)
	for _, v := range nums {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			unique = append(unique, v)
		}
	}
	sort.Ints(unique)
	m := len(unique)

	// build value -> index map
	comp := make(map[int]int, m)
	for i, v := range unique {
		comp[v] = i
	}

	// dp[i] = sum of products of magical sequences ending with unique[i]
	dp := make([]int64, m)
	for _, v := range nums {
		idx := comp[v]
		dp[idx] += int64(v)
	}
	if k == 1 {
		sum := int64(0)
		for _, val := range dp {
			sum += val
		}
		return int(sum % MOD)
	}

	// prefix[i] = prefix sum of dp[0..i-1]
	prefix := make([]int64, m+1)
	newDP := make([]int64, m) // reuse this allocation
	
	for length := 2; length <= k; length++ {
		// build prefix sums once per length
		for i := range m {
			prefix[i+1] = (prefix[i] + dp[i]) % MOD
		}

		// clear and reuse the slice
		for j := range m {
			v := int64(unique[j])
			// all indices < j contribute
			newDP[j] = (prefix[j] * v) % MOD
		}
		dp, newDP = newDP, dp // swap slices
	}

	sum := int64(0)
	for _, val := range dp {
		sum = (sum + val) % MOD
	}
	return int(sum)
}
