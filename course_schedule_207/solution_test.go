package course_schedule_207

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{"simple acyclic", 2, [][]int{{1, 0}}, true},
		{"simple cycle", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"3-cycle", 3, [][]int{{0, 1}, {1, 2}, {2, 0}}, false},
		{"acyclic three nodes", 3, [][]int{{0, 1}, {0, 2}, {1, 2}}, true},
		{"two prereqs", 3, [][]int{{0, 2}, {1, 0}}, true},
		{"single prereq", 3, [][]int{{0, 2}}, true},
		{"cycle via extra edge", 3, [][]int{{1, 0}, {2, 0}, {0, 2}}, false},
		{"no prereqs", 4, [][]int{}, true},
		{"empty graph", 0, [][]int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canFinish(tt.numCourses, tt.prerequisites))
			assert.Equal(t, tt.want, CanFinish(tt.numCourses, tt.prerequisites))
		})
	}
}
