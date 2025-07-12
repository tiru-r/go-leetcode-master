package longest_common_subsequence_1143

import (
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
