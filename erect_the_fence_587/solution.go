package erect_the_fence_587

import "sort"

// OuterTrees returns the vertices of the convex hull that contains every tree.
// It uses Andrew’s monotone-chain algorithm and runs in O(n log n).
func OuterTrees(trees [][]int) [][]int {
	n := len(trees)
	if n <= 1 {
		return trees
	}

	// 1. Sort lexicographically (x, then y).
	sort.Slice(trees, func(i, j int) bool {
		a, b := trees[i], trees[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] < b[1])
	})

	// 2. Monotone chain in one go.
	hull := make([][]int, 0, n)

	// Lower hull (left → right)
	for _, p := range trees {
		for len(hull) >= 2 && cross(hull[len(hull)-2], hull[len(hull)-1], p) < 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	// Upper hull (right → left) – skip the last point already in lower hull.
	upperStart := len(hull)
	for i := n - 2; i >= 0; i-- {
		p := trees[i]
		for len(hull) > upperStart && cross(hull[len(hull)-2], hull[len(hull)-1], p) < 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	// 3. Remove the duplicate first/last point.
	hull = hull[:len(hull)-1]
	return hull
}

// cross returns (a-o) × (b-o).
// >0 : ccw turn, <0 : cw turn, 0 : collinear.
func cross(o, a, b []int) int {
	return (a[0]-o[0])*(b[1]-o[1]) - (a[1]-o[1])*(b[0]-o[0])
}
