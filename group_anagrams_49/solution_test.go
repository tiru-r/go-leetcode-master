package group_anagrams_49

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	type Case struct {
		name  string
		input []string
		want  [][]string
	}

	cases := []Case{
		// 1. LeetCode official samples
		{"LC example 1",
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{"LC example 2", []string{""}, [][]string{{""}}},
		{"LC example 3", []string{"a"}, [][]string{{"a"}}},

		// 2. Empty / single
		{"empty slice", []string{}, [][]string{}},
		{"single string", []string{"abc"}, [][]string{{"abc"}}},

		// 3. Duplicates & all identical
		{"duplicate words", []string{"abc", "abc"}, [][]string{{"abc", "abc"}}},
		{"all identical", []string{"aaa", "aaa", "aaa"}, [][]string{{"aaa", "aaa", "aaa"}}},

		// 4. No anagrams
		{"all unique", []string{"a", "b", "c", "d"}, [][]string{{"a"}, {"b"}, {"c"}, {"d"}}},

		// 5. Long strings (26 unique letters)
		{"long anagram 26",
			[]string{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
			[][]string{{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}}},

		// 6. Stress – 100 identical anagrams
		{"stress 100 identical",
			repeat("abcdefghijklmnopqrstuvwxyz", 100),
			[][]string{repeat("abcdefghijklmnopqrstuvwxyz", 100)}},

		// 7. Mixed length & frequencies
		{"mixed lengths",
			[]string{"aabbcc", "abcabc", "abccba", "xyz", "zyx", "abcd"},
			[][]string{{"aabbcc", "abcabc", "abccba"}, {"xyz", "zyx"}, {"abcd"}}},

		// 8. Capital letters not present (constraint is lowercase)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := GroupAnagrams(tc.input)
			normalize(got)
			normalize(tc.want)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%s: got %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

/* ---------- helpers ---------- */

func normalize(groups [][]string) {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
}

func repeat(s string, n int) []string {
	out := make([]string, n)
	for i := range out {
		out[i] = s
	}
	return out
}
