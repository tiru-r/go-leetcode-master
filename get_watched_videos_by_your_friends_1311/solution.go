package get_watched_videos_by_your_friends_1311

import (
	"cmp"
	"slices"
)

func watchedVideosByFriends(
	watchedVideos [][]string,
	friends [][]int,
	id int,
	level int,
) []string {

	seen := map[int]bool{id: true}
	curr := []int{id}

	// BFS exactly `level` steps
	for level > 0 && len(curr) > 0 {
		next := []int{}
		for _, u := range curr {
			for _, v := range friends[u] {
				if !seen[v] {
					seen[v] = true
					next = append(next, v)
				}
			}
		}
		curr = next
		level--
	}

	// frequency map
	freq := map[string]int{}
	for _, uid := range curr {
		for _, vid := range watchedVideos[uid] {
			freq[vid]++
		}
	}

	// slice + stable sort (freq asc, then name asc)
	type pair struct {
		title string
		cnt   int
	}
	list := make([]pair, 0, len(freq))
	for t, c := range freq {
		list = append(list, pair{t, c})
	}
	slices.SortFunc(list, func(a, b pair) int {
		if c := cmp.Compare(a.cnt, b.cnt); c != 0 {
			return c
		}
		return cmp.Compare(a.title, b.title)
	})

	out := make([]string, len(list))
	for i, p := range list {
		out[i] = p.title
	}
	return out
}
