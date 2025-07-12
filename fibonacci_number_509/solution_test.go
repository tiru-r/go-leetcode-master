package fibonacci_number_509

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"fib 0", 0, 0},
		{"fib 1", 1, 1},
		{"fib 2", 2, 1},
		{"fib 3", 3, 2},
		{"fib 4", 4, 3},
		{"fib 5", 5, 5},
		{"fib 10", 10, 55},
		{"fib negative", -1, -1}, // optional edge-case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fib(tt.n))
		})
	}
}
