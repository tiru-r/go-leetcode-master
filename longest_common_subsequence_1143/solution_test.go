package longest_common_subsequence_1143

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{"empty strings", "", "", 0},
		{"one empty", "", "abc", 0},
		{"identical", "abc", "abc", 3},
		{"no overlap", "abc", "def", 0},
		{"single char match", "bsbininm", "jmjkbkjkv", 1},
		{"regular case", "abcde", "ace", 3},
		{"substring", "abc", "ab", 2},
		{"long strings", "pmjghexybyrgzczy", "hafcdqbgncrcbihk", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestCommonSubsequence(tt.text1, tt.text2))
		})
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	text1 := "abcdefghijklmnopqrstuvwxyz"
	text2 := "acegikmoqsuwy"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonSubsequence(text1, text2)
	}
}

func BenchmarkLongestCommonSubsequenceLarge(b *testing.B) {
	text1 := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 40)
	text2 := strings.Repeat("acegikmoqsuwy", 80)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonSubsequence(text1, text2)
	}
}
