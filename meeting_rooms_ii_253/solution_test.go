package meeting_rooms_ii_253

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMeetingRooms(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "single meeting",
			args: args{
				intervals: [][]int{
					{0, 30},
				},
			},
			want: 1,
		},
		{
			name: "overlapping meetings requiring 2 rooms",
			args: args{
				intervals: [][]int{
					{0, 30},
					{5, 10},
					{15, 20},
				},
			},
			want: 2,
		},
		{
			name: "no overlap, requiring 1 room",
			args: args{
				intervals: [][]int{
					{7, 10},
					{2, 4},
				},
			},
			want: 1,
		},
		{
			name: "multiple overlaps, requiring 2 rooms",
			args: args{
				intervals: [][]int{
					{9, 10},
					{4, 9},
					{4, 17},
				},
			},
			want: 2,
		},
		{
			name: "multiple meetings at same time, requiring 2 rooms",
			args: args{
				intervals: [][]int{
					{1, 5},
					{8, 9},
					{8, 9},
				},
			},
			want: 2,
		},
		{
			name: "empty intervals list",
			args: args{
				intervals: [][]int{},
			},
			want: 0,
		},
		{
			name: "multiple overlaps with varied times",
			args: args{
				intervals: [][]int{
					{1, 10},
					{2, 7},
					{3, 19},
					{8, 12},
					{10, 20},
				},
			},
			want: 3, // At time 3: [1,10], [2,7], [3,19] are active
		},
		{
			name: "long meeting covering others",
			args: args{
				intervals: [][]int{
					{1, 20},
					{2, 5},
					{6, 10},
				},
			},
			want: 2, // Correction: [1,20] overlaps with [2,5] and [6,10], requiring 2 rooms
		},
		{
			name: "meetings ending exactly when new ones start",
			args: args{
				intervals: [][]int{
					{1, 5},
					{5, 10},
					{10, 15},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minMeetingRooms(tt.args.intervals), "Number of rooms should match")
		})
	}
}
