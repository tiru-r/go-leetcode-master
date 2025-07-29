package critical_connections_in_a_network_1192

// criticalConnections returns all critical connections (bridges) in an undirected graph.
func criticalConnections(n int, connections [][]int) [][]int {
	// Build adjacency list with pre-allocated capacity
	adj := make([][]int, n)
	for i := range adj {
		adj[i] = make([]int, 0, 2) // Most nodes have 2-3 connections
	}
	for _, e := range connections {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	disc := make([]int, n)
	low := make([]int, n)
	visited := make([]bool, n)
	for i := range disc {
		disc[i] = -1
	}

	// Pre-allocate bridges slice - worst case is n-1 bridges (tree)
	bridges := make([][]int, 0, n-1)
	time := 0

	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		time++
		disc[u] = time
		low[u] = time
		visited[u] = true

		for _, v := range adj[u] {
			if v == parent {
				continue
			}

			if !visited[v] { // not visited
				dfs(v, u)
				low[u] = min(low[u], low[v])
				if low[v] > disc[u] {
					bridges = append(bridges, []int{u, v})
				}
			} else { // back edge
				low[u] = min(low[u], disc[v])
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, -1)
		}
	}

	return bridges
}

// CriticalConnections exists only for backward compatibility.
func CriticalConnections(n int, connections [][]int) [][]int {
	return criticalConnections(n, connections)
}
