package shortest_path_visiting_all_nodes_847

// State represents a node and the set of visited nodes
type State struct {
	node int
	mask int
}

// ShortestPathLength finds the shortest path that visits every node
// Time: O(n² * 2ⁿ), Space: O(n * 2ⁿ)
func ShortestPathLength(graph [][]int) int {
	n := len(graph)
	if n == 1 {
		return 0
	}
	
	target := (1 << n) - 1
	visited := make(map[State]bool)
	queue := make([]State, 0, n)
	
	// Start from all nodes
	for i := 0; i < n; i++ {
		state := State{node: i, mask: 1 << i}
		queue = append(queue, state)
		visited[state] = true
	}
	
	steps := 0
	
	for len(queue) > 0 {
		size := len(queue)
		
		for i := 0; i < size; i++ {
			curr := queue[i]
			
			if curr.mask == target {
				return steps
			}
			
			// Visit all neighbors
			for _, neighbor := range graph[curr.node] {
				newMask := curr.mask | (1 << neighbor)
				newState := State{node: neighbor, mask: newMask}
				
				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, newState)
				}
			}
		}
		
		queue = queue[size:]
		steps++
	}
	
	// If we reach here, the graph is disconnected
	// Return minimum steps to visit all components
	return n - 1
}
