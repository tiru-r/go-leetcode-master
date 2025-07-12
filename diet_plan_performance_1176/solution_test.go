package diet_plan_performance_1176

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDietPlanPerformance(t *testing.T) {
	tests := []struct {
		name            string
		calories        []int
		k, lower, upper int
		want            int
	}{
		{"single window exact bounds", []int{1, 2, 3, 4, 5}, 1, 3, 3, 0},
		{"two-day window over upper", []int{3, 2}, 2, 0, 1, 1},
		{"two-day window within bounds", []int{6, 5, 0, 0}, 2, 1, 5, 0},
		{"empty calories", []int{}, 2, 0, 0, 0},
		{"k larger than slice", []int{1, 2}, 5, 0, 10, 0},
		{"all below lower", []int{1, 1, 1}, 3, 4, 6, -1},
		{"all above upper", []int{10, 10}, 2, 0, 5, 1},
		{"mixed points", []int{1, 2, 3, 4, 5}, 3, 5, 7, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want,
				dietPlanPerformance(tt.calories, tt.k, tt.lower, tt.upper))
		})
	}
}
