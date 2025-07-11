package get_watched_videos_by_your_friends_1311

import (
	"cmp"
	"slices"
)

// Note: study again.
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	seen := map[int]bool{id: true}
	currentLevel := []int{id}

	// BFS for level times
	for level > 0 {
		var nextLevel []int
		for _, friendID := range currentLevel {
			for _, edge := range friends[friendID] {
				if !seen[edge] {
					nextLevel = append(nextLevel, edge)
					seen[edge] = true
				}
			}
		}
		currentLevel = nextLevel
		level--
	}

	// Count video frequencies
	freq := make(map[string]int)
	for _, friendID := range currentLevel {
		for _, video := range watchedVideos[friendID] {
			freq[video]++
		}
	}

	// Convert to slice for sorting
	type videoFreq struct {
		title string
		freq  int
	}

	videos := make([]videoFreq, 0, len(freq))
	for title, frequency := range freq {
		videos = append(videos, videoFreq{title, frequency})
	}

	// Sort by frequency, then alphabetically
	slices.SortFunc(videos, func(a, b videoFreq) int {
		if freqCmp := cmp.Compare(a.freq, b.freq); freqCmp != 0 {
			return freqCmp
		}
		return cmp.Compare(a.title, b.title)
	})

	// Extract titles
	result := make([]string, len(videos))
	for i, v := range videos {
		result[i] = v.title
	}

	return result
}
