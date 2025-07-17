package word_search_212

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindWords(t *testing.T) {
	tests := []struct {
		name   string
		board  [][]byte
		words  []string
		expect []string
	}{
		{"empty board", [][]byte{}, []string{"foo"}, []string{}},
		{"empty words", [][]byte{{'a'}}, []string{}, []string{}},
		{"single cell hit", [][]byte{{'a'}}, []string{"a"}, []string{"a"}},
		{"single cell miss", [][]byte{{'a'}}, []string{"b"}, []string{}},
		{"duplicates removed", [][]byte{
			{'a', 'a'},
		}, []string{"aa", "aa"}, []string{"aa"}}, // only once
		{"LeetCode 212 example",
			[][]byte{
				{'o', 'a', 'a', 'n'},
				{'e', 't', 'a', 'e'},
				{'i', 'h', 'k', 'r'},
				{'i', 'f', 'l', 'v'},
			},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"eat", "oath"},
		},
		{"prefix longer than board",
			[][]byte{{'a', 'b'}},
			[]string{"abc"},
			[]string{},
		},
		{"long snake",
			[][]byte{
				{'a', 'b', 'c', 'd', 'e'},
				{'f', 'g', 'h', 'i', 'j'},
				{'k', 'l', 'm', 'n', 'o'},
				{'p', 'q', 'r', 's', 't'},
			},
			[]string{"abcd", "abcde"},
			[]string{"abcd", "abcde"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findWords(tc.board, tc.words)
			sort.Strings(got)
			sort.Strings(tc.expect)
			if !reflect.DeepEqual(got, tc.expect) {
				t.Fatalf("got %v, want %v", got, tc.expect)
			}
		})
	}
}
