package lru_cache_146

import (
	"slices"
	"testing"
)

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
			cache := NewLRUCache(tt.capacity)
			results := []int{}

			for _, op := range tt.operations {
				switch op.op {
				case "put":
					cache.Put(op.key, op.value)
				case "get":
					results = append(results, cache.Get(op.key))
				}
			}

			if !slices.Equal(results, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, results)
			}
		})
	}
}

type operation struct {
	op    string // "put" or "get"
	key   int
	value int // used only for "put"
}

func TestNewLRUCache(t *testing.T) {
	c := NewLRUCache(2)
	if c.capacity != 2 {
		t.Errorf("expected capacity 2, got %d", c.capacity)
	}
	if len(c.cache) != 0 {
		t.Errorf("expected empty map, got len %d", len(c.cache))
	}
	if c.head == nil || c.tail == nil || c.head.next != c.tail || c.tail.prev != c.head {
		t.Error("sentinel links are wrong")
	}
}

func BenchmarkLRUCache(b *testing.B) {
	cache := NewLRUCache(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i%200, i)
		cache.Get(i % 100)
	}
}

func BenchmarkLRUCacheGet(b *testing.B) {
	cache := NewLRUCache(1000)
	for i := 0; i < 1000; i++ {
		cache.Put(i, i*10)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

func BenchmarkLRUCachePut(b *testing.B) {
	cache := NewLRUCache(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Put(i, i*10)
	}
}
