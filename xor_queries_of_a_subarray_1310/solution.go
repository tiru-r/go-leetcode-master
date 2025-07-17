package xor_queries_of_a_subarray_1310

func xorQueries(arr []int, queries [][]int) []int {
	if len(arr) == 0 || len(queries) == 0 {
		return []int{}
	}

	// Build prefix XOR array for O(1) range queries
	prefix := make([]int, len(arr)+1)
	for i, val := range arr {
		prefix[i+1] = prefix[i] ^ val
	}

	// Process queries using prefix XOR property: XOR(l,r) = prefix[r+1] ^ prefix[l]
	result := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		result[i] = prefix[r+1] ^ prefix[l]
	}
	return result
}
