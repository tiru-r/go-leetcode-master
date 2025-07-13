package counting_bits_338

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countBits(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{"n=2", 2, []int{0, 1, 1}},
		{"n=5", 5, []int{0, 1, 1, 2, 1, 2}},
		{"n=0", 0, []int{0}},
		{"n=1", 1, []int{0, 1}},
		{"n=8", 8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countBits(tt.n))
		})
	}
}
