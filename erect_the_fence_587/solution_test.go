package erect_the_fence_587

import "testing"

func TestOuterTrees(t *testing.T) {
	tests := []struct {
		trees [][]int
		want  [][]int
	}{
		{
			trees: [][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
			want:  [][]int{{1, 1}, {2, 0}, {4, 2}, {3, 3}, {2, 4}},
		},
		{
			trees: [][]int{{1, 2}, {2, 2}, {4, 2}},
			want:  [][]int{{1, 2}, {2, 2}, {4, 2}},
		},
		{
			trees: [][]int{{0, 0}},
			want:  [][]int{{0, 0}},
		},
		{
			trees: [][]int{{0, 0}, {1, 1}},
			want:  [][]int{{0, 0}, {1, 1}},
		},
		{
			trees: [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}},
			want:  [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}},
		},
		{
			trees: [][]int{{3, 0}, {4, 0}, {5, 0}, {6, 1}, {7, 2}, {7, 3}, {7, 4}, {6, 5}, {5, 5}, {4, 5}, {3, 5}, {2, 5}, {1, 4}, {1, 3}, {1, 2}, {2, 1}, {4, 2}, {0, 3}},
			want:  [][]int{{0, 3}, {1, 2}, {2, 1}, {3, 0}, {4, 0}, {5, 0}, {6, 1}, {7, 2}, {7, 3}, {7, 4}, {6, 5}, {5, 5}, {4, 5}, {3, 5}, {2, 5}, {1, 4}},
		},
	}

	for i, test := range tests {
		got := OuterTrees(test.trees)
		if !isValidConvexHull(got, test.want) {
			t.Errorf("Test %d: OuterTrees(%v) = %v; want %v", i, test.trees, got, test.want)
		}
	}
}

// isValidConvexHull checks if two convex hulls contain the same points
func isValidConvexHull(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}
	
	// Convert to sets for comparison
	gotSet := make(map[[2]int]bool)
	for _, p := range got {
		gotSet[[2]int{p[0], p[1]}] = true
	}
	
	wantSet := make(map[[2]int]bool)
	for _, p := range want {
		wantSet[[2]int{p[0], p[1]}] = true
	}
	
	// Check if sets are equal
	for p := range gotSet {
		if !wantSet[p] {
			return false
		}
	}
	
	for p := range wantSet {
		if !gotSet[p] {
			return false
		}
	}
	
	return true
}

func TestCross(t *testing.T) {
	tests := []struct {
		p1, p2, p3 []int
		want       int
	}{
		{[]int{0, 0}, []int{1, 0}, []int{1, 1}, 1},   // Counter-clockwise
		{[]int{0, 0}, []int{1, 0}, []int{1, -1}, -1}, // Clockwise
		{[]int{0, 0}, []int{1, 0}, []int{2, 0}, 0},   // Collinear
	}

	for i, test := range tests {
		got := cross(test.p1, test.p2, test.p3)
		if got != test.want {
			t.Errorf("Test %d: cross(%v, %v, %v) = %d; want %d", 
				i, test.p1, test.p2, test.p3, got, test.want)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Single point
	result := OuterTrees([][]int{{5, 5}})
	if len(result) != 1 || result[0][0] != 5 || result[0][1] != 5 {
		t.Errorf("Single point test failed: got %v", result)
	}
	
	// Two points
	result = OuterTrees([][]int{{0, 0}, {3, 4}})
	if len(result) != 2 {
		t.Errorf("Two points test failed: got %v", result)
	}
	
	// Collinear points
	result = OuterTrees([][]int{{0, 0}, {1, 1}, {2, 2}})
	if len(result) != 3 {
		t.Errorf("Collinear points test failed: got %v", result)
	}
}

func BenchmarkOuterTrees(b *testing.B) {
	// Generate test data
	trees := make([][]int, 100)
	for i := range trees {
		trees[i] = []int{i % 10, (i * 7) % 10}
	}
	
	b.ResetTimer()
	for range b.N {
		OuterTrees(trees)
	}
}

func BenchmarkOuterTreesLarge(b *testing.B) {
	// Generate larger test data
	trees := make([][]int, 1000)
	for i := range trees {
		trees[i] = []int{i % 100, (i * 17) % 100}
	}
	
	b.ResetTimer()
	for range b.N {
		OuterTrees(trees)
	}
}
