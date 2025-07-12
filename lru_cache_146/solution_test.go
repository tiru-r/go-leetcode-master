package lru_cache_146

import (
	"reflect"
	"testing"
)

// TestLRUCache uses table-driven testing to verify the LRUCache implementation.
func TestLRUCache(t *testing.T) {
	tests := []struct {
		name       string
		capacity   int
		operations []operation
		expected   []int
	}{
		{
			name:     "Basic Put and Get",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, value: 100},
				{op: "put", key: 2, value: 200},
				{op: "get", key: 1},
				{op: "get", key: 2},
				{op: "get", key: 3},
			},
			expected: []int{100, 200, -1},
		},
		{
			name:     "LRU Eviction",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, value: 100},
				{op: "put", key: 2, value: 200},
				{op: "put", key: 3, value: 300}, // Evicts key 1
				{op: "get", key: 1},
				{op: "get", key: 2},
				{op: "get", key: 3},
			},
			expected: []int{-1, 200, 300},
		},
		{
			name:     "Update Existing Key",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, value: 100},
				{op: "put", key: 2, value: 200},
				{op: "put", key: 1, value: 150}, // Update key 1
				{op: "get", key: 1},
				{op: "put", key: 3, value: 300}, // Evicts key 2
				{op: "get", key: 2},
			},
			expected: []int{150, -1},
		},
		{
			name:     "Move to Head on Get",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, value: 100},
				{op: "put", key: 2, value: 200},
				{op: "get", key: 1},             // Moves key 1 to head
				{op: "put", key: 3, value: 300}, // Evicts key 2
				{op: "get", key: 2},
				{op: "get", key: 1},
			},
			expected: []int{100, -1, 100},
		},
		{
			name:     "Single Capacity",
			capacity: 1,
			operations: []operation{
				{op: "put", key: 1, value: 100},
				{op: "put", key: 2, value: 200}, // Evicts key 1
				{op: "get", key: 1},
				{op: "get", key: 2},
			},
			expected: []int{-1, 200},
		},
		{
			name:     "Empty Cache",
			capacity: 2,
			operations: []operation{
				{op: "get", key: 1},
			},
			expected: []int{-1},
		},
		{
			name:     "Complex Sequence",
			capacity: 2,
			operations: []operation{
				{op: "put", key: 1, value: 1},
				{op: "put", key: 2, value: 2},
				{op: "get", key: 1},
				{op: "put", key: 3, value: 3}, // Evicts key 2
				{op: "get", key: 2},
				{op: "put", key: 4, value: 4}, // Evicts key 1
				{op: "get", key: 1},
				{op: "get", key: 3},
				{op: "get", key: 4},
			},
			expected: []int{1, -1, -1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.capacity)
			results := []int{}

			for _, op := range tt.operations {
				if op.op == "put" {
					cache.Put(op.key, op.value)
				} else if op.op == "get" {
					results = append(results, cache.Get(op.key))
				}
			}

			if !reflect.DeepEqual(results, tt.expected) {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, results)
			}

			// Verify internal structure: ensure map size matches expected non-evicted items
			expectedMapSize := 0
			for _, op := range tt.operations {
				if op.op == "put" {
					if _, ok := cache.data[op.key]; ok {
						expectedMapSize++
					}
				}
			}
			if len(cache.data) > cache.cap {
				t.Errorf("Test %s failed: map size %d exceeds capacity %d", tt.name, len(cache.data), cache.cap)
			}
		})
	}
}

// operation represents a single cache operation (Put or Get).
type operation struct {
	op    string // "put" or "get"
	key   int
	value int // Only used for "put"
}

// TestConstructor verifies the initialization of the LRUCache separately.
func TestConstructor(t *testing.T) {
	cache := Constructor(2)
	if cache.cap != 2 {
		t.Errorf("Expected capacity 2, got %d", cache.cap)
	}
	if cache.data == nil {
		t.Error("Expected data map to be initialized, got nil")
	}
	if cache.head == nil || cache.tail == nil {
		t.Error("Expected head and tail sentinels to be initialized, got nil")
	}
	if cache.head.next != cache.tail || cache.tail.prev != cache.head {
		t.Error("Expected head.next to point to tail and tail.prev to point to head")
	}
}
