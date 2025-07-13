package shortest_path_visiting_all_nodes_847

import "testing"

func TestShortestPathLength(t *testing.T) {
	tests := []struct {
		graph [][]int
		want  int
	}{
		{
			graph: [][]int{{1, 2, 3}, {0}, {0}, {0}},
			want:  4,
		},
		{
			graph: [][]int{{1}, {0, 2, 4}, {1, 3, 4}, {2}, {1, 2}},
			want:  4,
		},
		{
			graph: [][]int{{1, 2}, {0, 2}, {0, 1}},
			want:  2,
		},
		{
			graph: [][]int{{}},
			want:  0,
		},
		{
			graph: [][]int{{}, {}},
			want:  1,
		},
		{
			graph: [][]int{{1}, {0}},
			want:  1,
		},
		{
			graph: [][]int{{1, 2}, {0}, {0}},
			want:  2,
		},
		{
			graph: [][]int{{2, 3}, {2, 3}, {0, 1}, {0, 1}},
			want:  3,
		},
	}

	for i, test := range tests {
		got := ShortestPathLength(test.graph)
		if got != test.want {
			t.Errorf("Test %d: ShortestPathLength(%v) = %d; want %d", i, test.graph, got, test.want)
		}
	}
}

func TestState(t *testing.T) {
	s1 := State{node: 0, mask: 5}
	s2 := State{node: 0, mask: 5}
	s3 := State{node: 1, mask: 5}
	
	if s1 != s2 {
		t.Error("Equal states should be equal")
	}
	
	if s1 == s3 {
		t.Error("Different states should not be equal")
	}
}

func BenchmarkShortestPathLength(b *testing.B) {
	graph := [][]int{{1, 2, 3}, {0}, {0}, {0}}
	
	b.ResetTimer()
	for range b.N {
		ShortestPathLength(graph)
	}
}
