package longest_substring_without_repeating_characters_3

import (
	"strings"
	"testing"
)

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
		{"unicode", "pwwkewü", 4}, // 'ü' is a multi-byte rune - "kewü" = 4
		{"mixed ascii/unicode", "你好世界世", 4}, // "你好世界" = 4
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}

func BenchmarkLengthOfLongestSubstringASCII(b *testing.B) {
	s := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}

func BenchmarkLengthOfLongestSubstringUnicode(b *testing.B) {
	s := strings.Repeat("你好世界测试", 1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}
