package longest_repeating_character_replacement_424

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{"empty string", "", 0, 0},
		{"single char", "A", 0, 1},
		{"single char with replacements", "A", 1, 1},
		{"all same chars", "AAAAA", 0, 5},
		{"example 1", "AABABBA", 1, 4},
		{"full replacement", "BAAAB", 2, 5},
		{"replace all but one", "ABBB", 2, 4},
		{"alternating", "ABAB", 2, 4},
		{"longer sequence", "KRSCDCSONAJNHLBMDQGCJ", 4, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, characterReplacement(tt.s, tt.k))
		})
	}
}

func BenchmarkCharacterReplacement(b *testing.B) {
	s := "AABABBA"
	k := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		characterReplacement(s, k)
	}
}

func BenchmarkCharacterReplacementLarge(b *testing.B) {
	s := strings.Repeat("ABCDEFGHIJKLMNOP", 1000)
	k := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		characterReplacement(s, k)
	}
}
