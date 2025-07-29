package course_schedule_207

// canFinish determines whether it's possible to finish all courses.
// It uses Kahn's BFS-based topological sort.
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses <= 0 || len(prerequisites) == 0 {
		return true
	}

	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, p := range prerequisites {
		from, to := p[1], p[0] // p[1] -> p[0] : must take "from" before "to"
		adj[from] = append(adj[from], to)
		inDegree[to]++
	}

	q := make([]int, 0, numCourses)
	for i, d := range inDegree {
		if d == 0 {
			q = append(q, i)
		}
	}

	processed := 0
	head := 0
	for head < len(q) {
		c := q[head]
		head++
		processed++

		for _, nxt := range adj[c] {
			inDegree[nxt]--
			if inDegree[nxt] == 0 {
				q = append(q, nxt)
			}
		}
	}

	return processed == numCourses
}

// CanFinish is kept for backward-compatibility only.
func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}
