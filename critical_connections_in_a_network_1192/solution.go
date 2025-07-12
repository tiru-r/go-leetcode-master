package critical_connections_in_a_network_1192

// CanFinish returns all critical connections (bridges) in the graph.
// n   : number of vertices (0..n-1)
// con : undirected edges [u,v]
func CriticalConnections(n int, con [][]int) [][]int {
	// 1. Build adjacency list (slice version, one allocation)
	adj := make([][]int, n)
	for _, e := range con {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	// 2. Tarjan arrays
	disc := make([]int, n)
	low := make([]int, n)
	for i := range disc {
		disc[i], low[i] = -1, -1
	}

	var bridges [][]int
	time := 0

	// 3. DFS bridge finder
	var dfs func(int, int)
	dfs = func(u, parent int) {
		time++
		disc[u], low[u] = time, time
		for _, v := range adj[u] {
			if v == parent { // skip immediate parent
				continue
			}
			if disc[v] == -1 { // tree edge
				dfs(v, u)
				low[u] = min(low[u], low[v])
				if low[v] > disc[u] { // bridge detected
					bridges = append(bridges, []int{u, v})
				}
			} else { // back edge
				low[u] = min(low[u], disc[v])
			}
		}
	}

	// 4. Run DFS from every component (graph may be disconnected)
	for i := 0; i < n; i++ {
		if disc[i] == -1 {
			dfs(i, -1)
		}
	}
	return bridges
}
