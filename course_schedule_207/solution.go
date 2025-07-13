package course_schedule_207

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 {
		return true
	}
	
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	
	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		adj[from] = append(adj[from], to)
		inDegree[to]++
	}
	
	queue := make([]int, 0, numCourses)
	for i, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, i)
		}
	}
	
	processed := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		processed++
		
		for _, next := range adj[course] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	
	return processed == numCourses
}

// CanFinish is the original function name for backward compatibility
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}
