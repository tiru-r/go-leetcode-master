package letter_combinations_of_a_phone_number_17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{"empty", "", nil},
		{"single digit 4", "4", []string{"g", "h", "i"}},
		{"two digits 23", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"invalid digit", "12", nil},
		{"max depth 9", "9", []string{"w", "x", "y", "z"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, letterCombinations(tt.digits))
		})
	}
}
