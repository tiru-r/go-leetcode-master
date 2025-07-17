package xor_queries_of_a_subarray_1310

func xorQueries(arr []int, queries [][]int) []int {
	if len(arr) == 0 || len(queries) == 0 {
		return []int{}
	}

	// Build prefix-XOR array
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}

	// Answer queries
	res := make([]int, len(queries))
	for idx, q := range queries {
		l, r := q[0], q[1]
		// Defensive check (optional if inputs are guaranteed valid)
		if l < 0 || r >= n || l > r {
			res[idx] = 0
			continue
		}
		res[idx] = prefix[r+1] ^ prefix[l]
	}
	return res
}
