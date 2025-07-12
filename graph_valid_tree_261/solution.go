package graph_valid_tree_261

func validTree(n int, edges [][]int) bool {
	// 1. Edge count must be n-1 (acyclic + connected)
	if len(edges) != n-1 {
		return false
	}

	// 2. Handle trivial cases early
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	// 3. Build adjacency list
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 4. BFS/DFS to check connectivity
	seen := make([]bool, n)
	q := []int{0}
	seen[0] = true
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range adj[u] {
			if !seen[v] {
				seen[v] = true
				q = append(q, v)
			}
		}
	}

	// All nodes must be reachable
	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
