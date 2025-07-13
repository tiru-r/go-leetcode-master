package critical_connections_in_a_network_1192

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// normalize sorts every edge so that [u,v] and [v,u] compare equal.
func normalize(edges [][]int) [][]int {
	out := make([][]int, len(edges))
	for i, e := range edges {
		if e[0] > e[1] {
			out[i] = []int{e[1], e[0]}
		} else {
			out[i] = []int{e[0], e[1]}
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i][0] < out[j][0] ||
			(out[i][0] == out[j][0] && out[i][1] < out[j][1])
	})
	return out
}

func TestCriticalConnections(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		connections [][]int
		want        [][]int
	}{
		{"case 1", 4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}, [][]int{{1, 3}}},
		{"case 2", 6, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {3, 4}, {4, 5}, {5, 3}}, [][]int{{1, 3}}},
		{"case 3", 5, [][]int{{1, 0}, {2, 0}, {3, 2}, {4, 2}, {4, 3}, {3, 0}, {4, 0}}, [][]int{{1, 0}}},
		{"empty graph", 3, [][]int{}, [][]int{}},
		{"single node", 1, [][]int{}, [][]int{}},
		{"disjoint bridges", 4, [][]int{{0, 1}, {2, 3}}, [][]int{{0, 1}, {2, 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := normalize(criticalConnections(tt.n, tt.connections))
			got2 := normalize(CriticalConnections(tt.n, tt.connections))
			exp := normalize(tt.want)
			
			assert.Equal(t, exp, got1)
			assert.Equal(t, exp, got2)
		})
	}
}
