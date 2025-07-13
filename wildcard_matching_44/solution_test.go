package wildcard_matching_44

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
		{"adceb", "*a*b*", true},
		{"acdcb", "a*c?b", false},
		{"", "*", true},
		{"", "", true},
		{"a", "", false},
		{"", "a", false},
		{"abcdef", "a*f", true},
		{"abcdef", "a*d*f", true},
		{"abcdef", "a*d*g", false},
		{"mississippi", "m??*ss*?i*pi", false},
		{"hi", "*?", true},
		{"", "?", false},
		{"a", "?", true},
		{"ab", "?*", true},
		{"ab", "*?", true},
		{"abc", "a***c", true},
		{"abc", "a*****", true},
		{"leetcode", "*e*t*d*", true},
		{"character", "char", false},
		{"character", "*har*", true},
		{"character", "?haracter", true},
		{"character", "c?aracter", true}, // Corrected: ? matches 'h'
		{"ho", "ho**", true},
		{"babaaababaabababbbbbbaabaabbabababbaababbaaabbbaaab", "***bba**a*bbba**aab**b", false},
	}

	for i, test := range tests {
		got := IsMatch(test.s, test.p)
		if got != test.want {
			t.Errorf("Test %d: IsMatch(%q, %q) = %v; want %v", i, test.s, test.p, got, test.want)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	// Empty string and pattern
	if !IsMatch("", "") {
		t.Error("Empty string should match empty pattern")
	}
	
	// Only wildcards
	if !IsMatch("anything", "****") {
		t.Error("String should match pattern of only wildcards")
	}
	
	// Complex pattern
	if !IsMatch("abcdefg", "a*c*e*g") {
		t.Error("Should match complex wildcard pattern")
	}
}

func BenchmarkIsMatch(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	p := "a*m*z"
	
	b.ResetTimer()
	for range b.N {
		IsMatch(s, p)
	}
}

func BenchmarkIsMatchComplex(b *testing.B) {
	s := "babaaababaabababbbbbbaabaabbabababbaababbaaabbbaaab"
	p := "***bba**a*bbba**aab**b"
	
	b.ResetTimer()
	for range b.N {
		IsMatch(s, p)
	}
}
