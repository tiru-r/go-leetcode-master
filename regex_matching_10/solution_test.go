package regex_matching_10

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", "", true},
		{"", "a*", true},
		{"a", "", false},
		{"ab", ".*c", false},
		{"aaa", "a*a", true},
		{"aaa", "ab*a*c*a", true},
	}

	for i, tt := range tests {
		got := isMatch(tt.s, tt.p)
		if got != tt.want {
			t.Errorf("test %d: isMatch(%q, %q) = %v, want %v", 
				i, tt.s, tt.p, got, tt.want)
		}
	}
}
