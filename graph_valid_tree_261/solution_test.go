package graph_valid_tree_261

import (
	"testing"
)

func TestValidTree(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		// 1. LeetCode official samples
		{"LC sample 1", 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
		{"LC sample 2", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false},

		// 2. Trivial / boundary sizes
		{"n = 0", 0, [][]int{}, false},
		{"n = 1, no edges", 1, [][]int{}, true},
		{"n = 1, extra edge", 1, [][]int{{0, 0}}, false}, // self-loop ⇒ not n-1 edges
		{"n = 2, one edge", 2, [][]int{{0, 1}}, true},
		{"n = 2, no edge", 2, [][]int{}, false},
		{"n = 2, two edges", 2, [][]int{{0, 1}, {0, 1}}, false}, // duplicate edge

		// 3. Connectivity checks
		{"disconnected 4-nodes", 4, [][]int{{0, 1}, {2, 3}}, false},
		{"connected line 4-nodes", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}, true},
		{"star 5-nodes", 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, true},
		{"cycle 4-nodes", 4, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}}, false},

		// 4. Self-loops & multi-edges (should fail edge count)
		{"self-loop", 3, [][]int{{0, 1}, {1, 1}, {1, 2}}, false},
		{"multi-edge", 3, [][]int{{0, 1}, {0, 1}, {0, 2}}, false},

		// 5. Large graphs (performance sanity)
		{"large chain 1000", 1000, func() [][]int {
			e := make([][]int, 999)
			for i := 0; i < 999; i++ {
				e[i] = []int{i, i + 1}
			}
			return e
		}(), true},
		{"large star 1000", 1000, func() [][]int {
			e := make([][]int, 999)
			for i := 1; i < 1000; i++ {
				e[i-1] = []int{0, i}
			}
			return e
		}(), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validTree(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("%s: n=%d edges=%v ⇒ got %v, want %v",
					tt.name, tt.n, tt.edges, got, tt.want)
			}
		})
	}
}
