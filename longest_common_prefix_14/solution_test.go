package longest_common_prefix_14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"three strings with common prefix", []string{"flower", "flow", "flight"}, "fl"},
		{"no common prefix", []string{"dog", "racecar", "car"}, ""},
		{"single string", []string{"flower"}, "flower"},
		{"identical strings", []string{"abcd", "abcd"}, "abcd"},
		{"single character", []string{"a"}, "a"},
		{"empty slice", []string{}, ""},
		{"empty strings", []string{"", "b"}, ""},
		{"prefix is whole shortest", []string{"abc", "abcd", "abcde"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestCommonPrefix(tt.strs))
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{"flower", "flow", "flight"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(strs)
	}
}

func BenchmarkLongestCommonPrefixLarge(b *testing.B) {
	strs := make([]string, 1000)
	for i := range strs {
		strs[i] = "commonprefix" + string(rune('a'+i%26))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(strs)
	}
}
