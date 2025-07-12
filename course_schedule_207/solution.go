package course_schedule_207

// Note: study again.

// CanFinish reports whether all courses can be completed.
// It uses an iterative BFS (Kahn) for clarity and speed.
func CanFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return false
	}

	// 1. Build adjacency list and in-degree array.
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		from, to := p[1], p[0] // prerequisite -> course
		adj[from] = append(adj[from], to)
		inDegree[to]++
	}

	// 2. Seed queue with zero-in-degree courses.
	q := make([]int, 0, numCourses)
	for c, d := range inDegree {
		if d == 0 {
			q = append(q, c)
		}
	}

	// 3. Topological sort / cycle detection.
	seen := 0
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		seen++
		for _, nxt := range adj[v] {
			inDegree[nxt]--
			if inDegree[nxt] == 0 {
				q = append(q, nxt)
			}
		}
	}

	// If we processed every vertex, no cycle.
	return seen == numCourses
}
