package critical_connections_in_a_network_1192

func criticalConnections(n int, connections [][]int) [][]int {
	adj := make([][]int, n)
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	
	disc := make([]int, n)
	low := make([]int, n)
	for i := range disc {
		disc[i] = -1
		low[i] = -1
	}
	
	var bridges [][]int
	time := 0
	
	var dfs func(u, parent int)
	dfs = func(u, parent int) {
		time++
		disc[u] = time
		low[u] = time
		
		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			
			if disc[v] == -1 {
				dfs(v, u)
				low[u] = min(low[u], low[v])
				if low[v] > disc[u] {
					bridges = append(bridges, []int{u, v})
				}
			} else {
				low[u] = min(low[u], disc[v])
			}
		}
	}
	
	for i := range n {
		if disc[i] == -1 {
			dfs(i, -1)
		}
	}
	
	return bridges
}

// CriticalConnections is the original function name for backward compatibility
func CriticalConnections(n int, connections [][]int) [][]int {
	return criticalConnections(n, connections)
}
