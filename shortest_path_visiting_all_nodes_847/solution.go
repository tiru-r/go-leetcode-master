package shortest_path_visiting_all_nodes_847

// State represents a position in the BFS search space
type State struct{ node, visitedMask int }

// ShortestPathLength finds the shortest path visiting all nodes using BFS with bitmask.
// Time: O(n * 2^n), Space: O(n * 2^n) where n is the number of nodes.
func ShortestPathLength(graph [][]int) int {
	n := len(graph)
	if n == 1 {
		return 0
	}

	// Target state: all nodes visited (all bits set)
	targetMask := (1 << n) - 1
	
	// Use 2D array for O(1) visited lookups: visited[node][mask]
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, 1<<n)
	}
	
	// Efficient queue with pre-allocated capacity
	queue := make([]State, 0, n*n)
	queueHead := 0

	// Multi-source BFS: start from every node simultaneously for optimal paths
	for nodeIndex := range n {
		nodeMask := 1 << nodeIndex
		initialState := State{nodeIndex, nodeMask}
		queue = append(queue, initialState)
		visited[nodeIndex][nodeMask] = true
	}

	// BFS level-by-level processing
	for steps := 0; queueHead < len(queue); steps++ {
		levelSize := len(queue) - queueHead
		
		for range levelSize {
			currentState := queue[queueHead]
			queueHead++

			// Check if all nodes have been visited
			if currentState.visitedMask == targetMask {
				return steps
			}

			// Explore all neighbors of current node
			for _, neighborNode := range graph[currentState.node] {
				// Create new state by visiting the neighbor
				newMask := currentState.visitedMask | (1 << neighborNode)
				
				// Skip if this state has been visited before
				if !visited[neighborNode][newMask] {
					visited[neighborNode][newMask] = true
					nextState := State{neighborNode, newMask}
					queue = append(queue, nextState)
				}
			}
		}
	}
	// Should never reach here for valid connected graph
	return n - 1
}
