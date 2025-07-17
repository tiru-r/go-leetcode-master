package word_search_79

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name   string
		board  [][]byte
		word   string
		expect bool
	}{
		{"empty board", [][]byte{}, "A", false},
		{"empty word", [][]byte{{'A'}}, "", true},
		{"single char yes", [][]byte{{'A'}}, "A", true},
		{"single char no", [][]byte{{'A'}}, "B", false},
		{"2D single char", [][]byte{{'A', 'B'}}, "B", true},
		{"LeetCode example 1", [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCCED", true},
		{"LeetCode example 2", [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "SEE", true},
		{"LeetCode example 3", [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}, "ABCB", false},
		{"letter reuse attempt", [][]byte{
			{'A', 'A'},
		}, "AAA", false},
		{"backtracking required", [][]byte{
			{'A', 'B', 'C'},
			{'B', 'B', 'B'},
			{'C', 'B', 'A'},
		}, "ABBBBA", true},
		{"large word through long path", [][]byte{
			{'A', 'B', 'C', 'D'},
			{'E', 'F', 'G', 'H'},
			{'I', 'J', 'K', 'L'},
			{'M', 'N', 'O', 'P'},
		}, "AFKLOP", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := exist(tc.board, tc.word)
			if got != tc.expect {
				t.Errorf("exist(%q, %q) = %v; want %v",
					stringify(tc.board), tc.word, got, tc.expect)
			}
		})
	}
}

// stringify is just for nicer error messages
func stringify(b [][]byte) string {
	s := "["
	for i, row := range b {
		if i > 0 {
			s += " "
		}
		s += string(row)
		if i < len(b)-1 {
			s += "\n"
		}
	}
	s += "]"
	return s
}
