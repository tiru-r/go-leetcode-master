package graph_valid_tree_261

// validTree checks if an undirected graph forms a valid tree
// Optimized: O(V+E) time, O(V) space using Union-Find with path compression
func validTree(n int, edges [][]int) bool {
	// Tree property: exactly n-1 edges for n nodes
	if len(edges) != n-1 {
		return false
	}
	
	// Edge cases
	if n <= 1 {
		return n == 1 // n=0 is invalid, n=1 is valid
	}
	
	// Union-Find with path compression for efficient cycle detection
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}
	
	// Find with path compression
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x]) // Path compression
		}
		return parent[x]
	}
	
	// Union operation with connectivity tracking
	components := n
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		rootU, rootV := find(u), find(v)
		
		// If already connected, we have a cycle
		if rootU == rootV {
			return false
		}
		
		// Union the components
		parent[rootU] = rootV
		components--
	}
	
	// Valid tree: connected (exactly 1 component) and acyclic
	return components == 1
}
