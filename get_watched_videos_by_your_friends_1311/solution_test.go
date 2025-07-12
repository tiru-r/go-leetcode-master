package get_watched_videos_by_your_friends_1311

import (
	"reflect"
	"testing"
)

func TestWatchedVideosByFriends(t *testing.T) {
	type Case struct {
		name          string
		watchedVideos [][]string
		friends       [][]int
		id, level     int
		want          []string
	}

	cases := []Case{
		// 1. LeetCode official samples
		{"LC example 1",
			[][]string{{"A", "A"}, {"B", "B"}, {"C"}, {"D", "D"}},
			[][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
			0, 1,
			[]string{"B", "C"}},
		{"LC example 2",
			[][]string{{"A", "A"}, {"B", "B"}, {"C"}, {"D", "D"}},
			[][]int{{1, 2}, {0, 3}, {0, 3}, {1, 2}},
			0, 2,
			[]string{"D"}},

		// 2. Edge cases
		{"level 0 only self",
			[][]string{{"X", "Y"}},
			[][]int{{}},
			0, 0,
			[]string{"X", "Y"}},
		{"level unreachable (empty friends)",
			[][]string{{"A"}},
			[][]int{{}},
			0, 5,
			[]string{}},
		{"level unreachable (no edges)",
			[][]string{{}, {}, {}},
			[][]int{{}, {}, {}},
			0, 1,
			[]string{}},
		{"duplicate videos same frequency",
			[][]string{{}, {"A", "A"}, {"A", "B"}},
			[][]int{{1, 2}, {0}, {0}},
			0, 1,
			[]string{"A", "B"}},
		{"alphabetical tie-break",
			[][]string{{}, {"Z", "A"}, {"Z", "B"}},
			[][]int{{1, 2}, {0}, {0}},
			0, 1,
			[]string{"A", "B", "Z"}},
		{"all same video",
			[][]string{{}, {"X"}, {"X"}},
			[][]int{{1, 2}, {0}, {0}},
			0, 1,
			[]string{"X"}},

		// 3. Larger / stress
		{"chain length 4",
			[][]string{{"L0"}, {"L1"}, {"L2"}, {"L3"}, {"L4"}},
			[][]int{{1}, {2}, {3}, {4}, {}},
			0, 4,
			[]string{"L4"}},
		{"star graph level 2",
			[][]string{{"C"}, {"X"}, {"Y"}, {"Z"}},
			[][]int{{1, 2, 3}, {0}, {0}, {0}},
			0, 2,
			[]string{}}, // no friends-of-friends
		{"dense 5-node level 2",
			[][]string{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}},
			[][]int{{1, 2}, {0, 2, 3}, {0, 1, 4}, {1, 4}, {2, 3}},
			0, 2,
			[]string{"C", "D", "E"}}, // alphabetical order
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := watchedVideosByFriends(c.watchedVideos, c.friends, c.id, c.level)
			if !reflect.DeepEqual(got, c.want) {
				t.Errorf("%s:\ngot  %v\nwant %v", c.name, got, c.want)
			}
		})
	}
}
