package reverse_words_in_a_string_151

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name, s, want string
	}{
		{"basic", "the sky is blue", "blue is sky the"},
		{"multiple spaces", "  hello   world!  ", "world! hello"},
		{"single word", "word", "word"},
		{"all spaces", "     ", ""},
		{"empty", "", ""},
		{"unicode", "  你好 世界  ", "世界 你好"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverseWords(tt.s))
		})
	}
}
