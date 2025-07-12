package distribute_candies_575

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		name    string
		candies []int
		want    int
	}{
		{"half unique", []int{1, 1, 2, 2, 3, 3}, 3},
		{"unique capped by half", []int{1, 1, 2, 3}, 2},
		{"unique exceeds half", []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3}, 3},
		{"large values", []int{1000, 1000, 2, 1, 2, 5, 3, 1}, 4},
		{"single type", []int{5, 5, 5, 5}, 2},
		{"all unique", []int{1, 2, 3, 4}, 2},
		{"empty", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, distributeCandies(tt.candies))
		})
	}
}
