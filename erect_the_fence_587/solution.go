package erect_the_fence_587

import (
	"slices"
)

// Point represents a 2D point
type Point [2]int

// OuterTrees finds all points on the convex hull boundary (fence)
// Uses Andrew's monotone chain algorithm - more robust than Graham scan
// Time: O(n log n), Space: O(n)
func OuterTrees(trees [][]int) [][]int {
	n := len(trees)
	if n <= 1 {
		return trees
	}
	
	// Convert to Point slice for easier handling
	points := make([]Point, n)
	for i, tree := range trees {
		points[i] = Point{tree[0], tree[1]}
	}
	
	// Sort points lexicographically (x, then y)
	slices.SortFunc(points, func(a, b Point) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
	
	// Build lower hull
	lower := buildHull(points, false)
	
	// Build upper hull
	upper := buildHull(points, true)
	
	// Combine hulls and remove duplicates
	hullSet := make(map[Point]bool)
	
	// Add lower hull points
	for _, p := range lower {
		hullSet[p] = true
	}
	
	// Add upper hull points
	for _, p := range upper {
		hullSet[p] = true
	}
	
	// Convert back to result format
	result := make([][]int, 0, len(hullSet))
	for p := range hullSet {
		result = append(result, []int{p[0], p[1]})
	}
	
	return result
}

// buildHull constructs half of the convex hull
func buildHull(points []Point, reverse bool) []Point {
	pts := points
	if reverse {
		pts = make([]Point, len(points))
		copy(pts, points)
		slices.Reverse(pts)
	}
	
	hull := make([]Point, 0, len(pts))
	
	for _, p := range pts {
		// Remove points that create right turn (concave)
		for len(hull) >= 2 && crossProduct(hull[len(hull)-2], hull[len(hull)-1], p) < 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}
	
	return hull
}

// crossProduct calculates cross product of vectors (p1->p2) and (p1->p3)
// Returns positive for counter-clockwise turn, negative for clockwise
func crossProduct(p1, p2, p3 Point) int {
	return (p2[0]-p1[0])*(p3[1]-p1[1]) - (p2[1]-p1[1])*(p3[0]-p1[0])
}
