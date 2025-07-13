package erect_the_fence_587

import "slices"

func OuterTrees(trees [][]int) [][]int {
	n := len(trees)
	if n <= 3 {
		return trees
	}
	
	slices.SortFunc(trees, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
	
	hull := convexHull(trees)
	seen := make(map[[2]int]struct{}, len(hull))
	
	for _, p := range hull {
		seen[[2]int{p[0], p[1]}] = struct{}{}
	}
	
	result := make([][]int, 0, len(seen))
	for p := range seen {
		result = append(result, []int{p[0], p[1]})
	}
	
	return result
}

func convexHull(points [][]int) [][]int {
	lower := monotoneChain(points)
	slices.Reverse(points)
	upper := monotoneChain(points)
	
	return append(lower, upper...)
}

func monotoneChain(points [][]int) [][]int {
	hull := make([][]int, 0, len(points))
	
	for _, p := range points {
		for len(hull) >= 2 && cross(hull[len(hull)-2], hull[len(hull)-1], p) < 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}
	
	return hull
}

func cross(o, a, b []int) int {
	return (a[0]-o[0])*(b[1]-o[1]) - (a[1]-o[1])*(b[0]-o[0])
}
