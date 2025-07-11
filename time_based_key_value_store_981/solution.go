package time_based_key_value_store_981

import (
	"cmp"
	"slices"
)

type entry struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]*entry
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]*entry),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	entries := m.m[key]
	newEntry := &entry{value: value, timestamp: timestamp}

	// Find insertion point using binary search
	index, found := slices.BinarySearchFunc(entries, newEntry, func(a, b *entry) int {
		return cmp.Compare(a.timestamp, b.timestamp)
	})

	if found {
		// Replace existing entry with same timestamp
		entries[index] = newEntry
	} else {
		// Insert at the correct position
		entries = slices.Insert(entries, index, newEntry)
	}

	m.m[key] = entries
}

func (m *TimeMap) Get(key string, timestamp int) string {
	entries, ok := m.m[key]
	if !ok {
		return ""
	}

	// Binary search for the largest timestamp <= given timestamp
	target := &entry{timestamp: timestamp}
	index, found := slices.BinarySearchFunc(entries, target, func(a, b *entry) int {
		return cmp.Compare(a.timestamp, b.timestamp)
	})

	if found {
		return entries[index].value
	}
	if index == 0 {
		return ""
	}
	return entries[index-1].value
}
