package shortest_path_visiting_all_nodes_847

type State struct {
	node, mask int
}

func ShortestPathLength(graph [][]int) int {
	n := len(graph)
	if n == 1 {
		return 0
	}
	
	target := (1 << n) - 1
	visited := make(map[State]bool)
	queue := make([]State, 0, n)
	
	for i := range n {
		state := State{i, 1 << i}
		queue = append(queue, state)
		visited[state] = true
	}
	
	for steps := 0; len(queue) > 0; steps++ {
		size := len(queue)
		
		for range size {
			curr := queue[0]
			queue = queue[1:]
			
			if curr.mask == target {
				return steps
			}
			
			for _, neighbor := range graph[curr.node] {
				newState := State{neighbor, curr.mask | (1 << neighbor)}
				
				if !visited[newState] {
					visited[newState] = true
					queue = append(queue, newState)
				}
			}
		}
	}
	
	return n - 1
}
