package time_based_key_value_store_981

import (
	"cmp"
	"slices"
)

type timeEntry struct {
	value     string
	timestamp int
}

type TimeMap struct {
	data map[string][]timeEntry
}

func NewTimeMap() *TimeMap {
	return &TimeMap{data: make(map[string][]timeEntry)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	entry := timeEntry{value: value, timestamp: timestamp}
	entries := tm.data[key]

	index, found := slices.BinarySearchFunc(entries, entry,
		func(a, b timeEntry) int { return cmp.Compare(a.timestamp, b.timestamp) })

	if found {
		entries[index] = entry
	} else {
		entries = slices.Insert(entries, index, entry)
	}
	tm.data[key] = entries
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	entries, exists := tm.data[key]
	if !exists {
		return ""
	}

	entry := timeEntry{timestamp: timestamp}
	index, exact := slices.BinarySearchFunc(entries, entry,
		func(a, b timeEntry) int { return cmp.Compare(a.timestamp, b.timestamp) })

	if exact {
		return entries[index].value
	}
	if index == 0 {
		return ""
	}
	return entries[index-1].value
}
