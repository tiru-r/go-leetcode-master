package implement_trie_208

import "testing"

func TestTrie(t *testing.T) {
	tests := []struct {
		name  string
		ops   []string // insert, search, startsWith
		args  []string // arguments for each op
		wantB []bool   // expected bool results
	}{
		{
			name:  "basic ops",
			ops:   []string{"insert", "search", "search", "startsWith", "insert", "search"},
			args:  []string{"apple", "apple", "app", "app", "app", "app"},
			wantB: []bool{true, false, true, true},
		},
		{
			name:  "empty string",
			ops:   []string{"insert", "search", "startsWith"},
			args:  []string{"", "", ""},
			wantB: []bool{true, true, true},
		},
		{
			name:  "overwrite same word",
			ops:   []string{"insert", "insert", "search"},
			args:  []string{"abc", "abc", "abc"},
			wantB: []bool{true, true, true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tr := NewTrie()
			out := 0
			for i, op := range tc.ops {
				switch op {
				case "insert":
					tr.Insert(tc.args[i])
				case "search":
					got := tr.Search(tc.args[i])
					if got != tc.wantB[out] {
						t.Fatalf("Search(%q): want %v, got %v", tc.args[i], tc.wantB[out], got)
					}
					out++
				case "startsWith":
					got := tr.StartsWith(tc.args[i])
					if got != tc.wantB[out] {
						t.Fatalf("StartsWith(%q): want %v, got %v", tc.args[i], tc.wantB[out], got)
					}
					out++
				}
			}
		})
	}
}

func BenchmarkInsert(b *testing.B) {
	tr := NewTrie()
	for range b.N {
		tr.Insert("benchmark")
	}
}

func BenchmarkSearch(b *testing.B) {
	tr := NewTrie()
	tr.Insert("benchmark")
	b.ResetTimer()
	for range b.N {
		_ = tr.Search("benchmark")
	}
}
