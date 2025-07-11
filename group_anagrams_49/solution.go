package group_anagrams_49

import (
	"maps"
	"slices"
	"strings"
)

// Third solution based on creating identical encoding for
// strings that are anagrams of each other for grouping.
func groupAnagrams(strs []string) [][]string {
	groupMap := make(map[string][]string)
	for _, s := range strs {
		se := encodeString(s)
		groupMap[se] = append(groupMap[se], s)
	}

	// Use slices.Collect with maps.Values() to extract all groups efficiently
	return slices.Collect(maps.Values(groupMap))
}

func encodeString(s string) string {
	chars := make([]int, 26)
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
	}

	var encoded strings.Builder
	encoded.Grow(len(s) * 2) // Better capacity estimation
	for i, count := range chars {
		if count > 0 {
			encoded.WriteByte(byte('a' + i))
			encoded.WriteByte(byte('0' + count))
		}
	}

	return encoded.String()
}

// Second solution based on sorting the strings
func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		sorted := sortString(s)
		m[sorted] = append(m[sorted], s)
	}

	// Use slices.Collect with maps.Values() for efficient value extraction
	return slices.Collect(maps.Values(m))
}

func sortString(s string) string {
	chars := []rune(s)
	slices.Sort(chars)
	return string(chars)
}

// Note: Wrong approach. Good lessons though!
func groupAnagrams0(strs []string) [][]string {
	groups := make([][]string, 0)
	seen := make(map[string]bool)
	for i := 0; i < len(strs); i++ {
		_, ok := seen[strs[i]]
		if ok && strs[i] != "" {
			continue
		}
		group := make([]string, 1)
		group[0] = strs[i]
		for j := i + 1; j < len(strs); j++ {
			_, ok := seen[strs[j]]
			if !ok && isAnagram(strs[i], strs[j]) {
				group = append(group, strs[j])
				seen[strs[j]] = true
			}
		}
		groups = append(groups, group)
	}
	return groups
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if len(s1) == 0 {
		return true
	}

	m := make(map[rune]int)
	for _, ch := range s1 {
		m[ch]++
	}

	for _, ch := range s2 {
		if m[ch] == 0 {
			return false
		}
		m[ch]--
	}

	return true
}
