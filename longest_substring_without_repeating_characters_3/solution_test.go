package longest_substring_without_repeating_characters_3

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"single", "a", 1},
		{"all unique", "abcdef", 6},
		{"all same", "aaaaa", 1},
		{"duplicate at end", "abca", 3},
		{"duplicate at start", "aabcd", 4},
		{"duplicate in middle", "abcabcbb", 3},
		{"window jump twice", "abba", 2},
		{"unicode", "pwwkewü", 6}, // 'ü' is a multi-byte rune
		{"mixed ascii/unicode", "你好世界世", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

// Benchmarks the implementation on a worst-case input (no repeats).
func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	const n = 1_000
	s := make([]rune, n)
	for i := 0; i < n; i++ {
		s[i] = rune('a' + (i % 26))
	}
	in := string(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lengthOfLongestSubstring(in)
	}
}
