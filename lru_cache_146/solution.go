package lru_cache_146

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	tail     *node
}

type node struct {
	key, value int
	prev, next *node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	node, exists := c.cache[key]
	if !exists {
		return -1
	}
	c.moveToHead(node)
	return node.value
}

func (c *LRUCache) Put(key, value int) {
	if node, exists := c.cache[key]; exists {
		node.value = value
		c.moveToHead(node)
		return
	}

	node := &node{key: key, value: value}
	
	if len(c.cache) >= c.capacity {
		lru := c.tail.prev
		c.removeNode(lru)
		delete(c.cache, lru.key)
	}

	c.cache[key] = node
	c.addToHead(node)
}

func (c *LRUCache) moveToHead(node *node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) addToHead(node *node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) removeNode(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
