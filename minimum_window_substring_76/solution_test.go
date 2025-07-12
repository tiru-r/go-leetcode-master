package minimum_window_substring_76

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name, s, t, want string
	}{
		{"duplicate chars", "ABAACBAB", "ABC", "ACB"},
		{"classic leetcode", "ZADOBECOAEBANC", "ABC", "BANC"},
		{"single exact", "A", "A", "A"},
		{"single missing", "A", "B", ""},
		{"two chars target first", "AB", "A", "A"},
		{"two chars target last", "AB", "B", "B"},
		{"duplicates needed", "aa", "aa", "aa"},
		{"impossible", "a", "aa", ""},
		{"empty s", "", "a", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minWindow(tt.s, tt.t))
		})
	}
}
