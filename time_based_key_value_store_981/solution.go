// Package time_based_key_value_store_981 solves LeetCode 981.
package time_based_key_value_store_981

import (
	"cmp"
	"slices"
)

type entry struct {
	value     string
	timestamp int
}

// TimeMap maps keys to a chronologically-sorted slice of entries.
type TimeMap struct {
	m map[string][]entry
}

// Constructor returns an empty store.
func Constructor() TimeMap {
	return TimeMap{m: make(map[string][]entry)}
}

// Set inserts or overwrites the value at the given key and timestamp.
// Time: O(log N) for the binary search + O(N) for the slice insert.
// Space: O(1) amortised.
func (tm *TimeMap) Set(key string, value string, timestamp int) {
	e := entry{value: value, timestamp: timestamp}
	ents := tm.m[key]

	// Binary search to find the first entry with timestamp >= target.
	idx, found := slices.BinarySearchFunc(ents, e,
		func(a, b entry) int { return cmp.Compare(a.timestamp, b.timestamp) })

	if found { // overwrite existing timestamp
		ents[idx] = e
	} else { // insert in order
		ents = slices.Insert(ents, idx, e)
	}
	tm.m[key] = ents
}

// Get returns the most recent value for key whose timestamp <= target.
// If no such timestamp exists, returns "".
// Time: O(log N) via binary search.
func (tm *TimeMap) Get(key string, timestamp int) string {
	ents, ok := tm.m[key]
	if !ok {
		return ""
	}

	e := entry{timestamp: timestamp}
	idx, exact := slices.BinarySearchFunc(ents, e,
		func(a, b entry) int { return cmp.Compare(a.timestamp, b.timestamp) })

	if exact {
		return ents[idx].value
	}
	if idx == 0 {
		return ""
	}
	return ents[idx-1].value
}
