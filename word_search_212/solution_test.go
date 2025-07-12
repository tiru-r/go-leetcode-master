package word_search_212

import (
    "reflect"
    "sort"
    "testing"
)

func TestFindWords(t *testing.T) {
    tests := []struct {
        name     string
        board    [][]byte
        words    []string
        expected []string
    }{
        {
            name: "Example 1",
            board: [][]byte{
                {'o', 'a', 'a', 'n'},
                {'e', 't', 'a', 'e'},
                {'i', 'h', 'k', 'r'},
                {'i', 'f', 'l', 'v'},
            },
            words:    []string{"oath", "pea", "eat", "rain"},
            expected: []string{"eat", "oath"},
        },
        {
            name: "Example 2",
            board: [][]byte{
                {'a', 'b'},
                {'c', 'd'},
            },
            words:    []string{"abcb"},
            expected: []string{},
        },
        {
            name: "Single character",
            board: [][]byte{
                {'a'},
            },
            words:    []string{"a"},
            expected: []string{"a"},
        },
        {
            name: "Multiple words found",
            board: [][]byte{
                {'a', 'b', 'c'},
                {'d', 'e', 'f'},
                {'g', 'h', 'i'},
            },
            words:    []string{"abc", "adg", "ghi", "xyz"},
            expected: []string{"abc", "adg", "ghi"},
        },
        {
            name:     "Empty board",
            board:    [][]byte{},
            words:    []string{"test"},
            expected: []string{},
        },
        {
            name: "Empty words",
            board: [][]byte{
                {'a', 'b'},
                {'c', 'd'},
            },
            words:    []string{},
            expected: []string{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := findWords(tt.board, tt.words)
            if result == nil {
                result = []string{}
            }
            if tt.expected == nil {
                tt.expected = []string{}
            }
            sort.Strings(result)
            sort.Strings(tt.expected)
            
            if !reflect.DeepEqual(result, tt.expected) {
                t.Errorf("findWords() = %v, want %v", result, tt.expected)
            }
        })
    }
}
