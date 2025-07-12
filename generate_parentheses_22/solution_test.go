package generate_parentheses_22

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{0, []string{""}},
		{1, []string{"()"}},
		{2, []string{"()()", "(())"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{4, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}},
		{5, []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"}},
		{6, nil}, // too long to inline, checked by length
		{7, nil},
		{8, nil},
	}

	// helper to check length only for large n
	catalan := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430}

	for _, tt := range tests {
		got := generateParenthesis(tt.n)
		if tt.n >= 6 {
			if len(got) != catalan[tt.n] {
				t.Errorf("n=%d: got %d results, want %d", tt.n, len(got), catalan[tt.n])
			}
			continue
		}
		if !equal(got, tt.want) {
			t.Errorf("n=%d: got %v, want %v", tt.n, got, tt.want)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]int, len(a))
	for _, s := range a {
		m[s]++
	}
	for _, s := range b {
		if m[s] == 0 {
			return false
		}
		m[s]--
	}
	return true
}
