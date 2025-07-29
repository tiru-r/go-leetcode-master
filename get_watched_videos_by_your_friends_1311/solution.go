package get_watched_videos_by_your_friends_1311

import (
	"cmp"
	"slices"
)

// watchedVideosByFriends returns the videos watched by friends at
// the exact distance `level`, sorted by frequency ascending,
// then lexicographically ascending.
func watchedVideosByFriends(
	watchedVideos [][]string,
	friends [][]int,
	id, level int,
) []string {

	// ---------- special-case level 0 ----------
	if level == 0 {
		// unique videos of the user, alphabetical
		set := make(map[string]struct{})
		for _, v := range watchedVideos[id] {
			set[v] = struct{}{}
		}
		out := make([]string, 0, len(set))
		for v := range set {
			out = append(out, v)
		}
		slices.Sort(out)
		return out
	}

	// BFS to level k with efficient queue
	visited := make([]bool, len(friends))
	q := make([]int, 0, len(friends)) // pre-allocate with max capacity
	q = append(q, id)
	visited[id] = true
	head := 0

	for lvl := 0; lvl < level && head < len(q); lvl++ {
		tail := len(q)
		for head < tail {
			u := q[head]
			head++
			for _, v := range friends[u] {
				if !visited[v] {
					visited[v] = true
					q = append(q, v)
				}
			}
		}
	}

	// Aggregate videos from friends at target level
	freq := make(map[string]int)
	globalSeen := make(map[string]map[int]struct{}) // video -> set of users who watched it
	
	// Get friends at target level (remaining items in queue after BFS)
	for i := head; i < len(q); i++ {
		u := q[i]
		for _, video := range watchedVideos[u] {
			if globalSeen[video] == nil {
				globalSeen[video] = make(map[int]struct{})
			}
			globalSeen[video][u] = struct{}{}
		}
	}
	
	// Count unique friends per video
	for video, users := range globalSeen {
		freq[video] = len(users)
	}

	// Sort by frequency (ascending), then lexicographically
	keys := make([]string, 0, len(freq))
	for v := range freq {
		keys = append(keys, v)
	}
	slices.SortFunc(keys, func(a, b string) int {
		// Cache frequency lookups
		freqA, freqB := freq[a], freq[b]
		if freqA != freqB {
			return cmp.Compare(freqA, freqB)
		}
		return cmp.Compare(a, b)
	})
	return keys
}
