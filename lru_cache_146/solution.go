package lru_cache_146

// LRUCache implements a Least Recently Used cache with O(1) operations.
type LRUCache struct {
	capacity int
	data     map[int]*entry
	head     *entry // dummy head (MRU is head.next)
}

type entry struct {
	key, val   int
	prev, next *entry
}

// NewLRUCache creates a new LRU cache with the specified capacity.
func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("LRU cache capacity must be positive")
	}
	head := &entry{}
	head.prev, head.next = head, head // circular sentinel
	return &LRUCache{
		capacity: capacity,
		data:     make(map[int]*entry, capacity),
		head:     head,
	}
}

// moveToFront moves the given entry to the front of the list (most recently used).
func (c *LRUCache) moveToFront(node *entry) {
	// Unlink from current position
	node.prev.next, node.next.prev = node.next, node.prev
	// Insert after head (most recent position)
	node.prev, node.next = c.head, c.head.next
	c.head.next.prev, c.head.next = node, node
}

// Get retrieves a value from the cache and marks it as recently used.
func (c *LRUCache) Get(key int) int {
	if node, exists := c.data[key]; exists {
		c.moveToFront(node)
		return node.val
	}
	return -1
}

// Put adds or updates a key-value pair in the cache.
func (c *LRUCache) Put(key, value int) {
	if node, exists := c.data[key]; exists {
		// Update existing entry and move to front
		node.val = value
		c.moveToFront(node)
		return
	}

	// Check if we need to evict the least recently used item
	if len(c.data) == c.capacity {
		// Remove LRU (node just before head in circular list)
		lru := c.head.prev
		lru.prev.next, lru.next.prev = lru.next, lru.prev
		delete(c.data, lru.key)
	}

	// Create and insert new node at front
	newNode := &entry{key: key, val: value}
	newNode.prev, newNode.next = c.head, c.head.next
	c.head.next.prev, c.head.next = newNode, newNode
	c.data[key] = newNode
}
