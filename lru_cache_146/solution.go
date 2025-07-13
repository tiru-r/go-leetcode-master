package lru_cache_146

type LRUCache struct {
	capacity int
	cache    map[int]*cacheNode
	head     *cacheNode
	tail     *cacheNode
}

type cacheNode struct {
	prev  *cacheNode
	next  *cacheNode
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	head, tail := &cacheNode{}, &cacheNode{}
	head.next, tail.prev = tail, head

	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*cacheNode, capacity),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if node := c.cache[key]; node != nil {
		c.moveToHead(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if node := c.cache[key]; node != nil {
		node.value = value
		c.moveToHead(node)
		return
	}

	if len(c.cache) == c.capacity {
		lru := c.tail.prev
		delete(c.cache, lru.key)
		c.unlink(lru)
	}

	node := &cacheNode{key: key, value: value}
	c.cache[key] = node
	c.linkHead(node)
}

func (c *LRUCache) moveToHead(node *cacheNode) {
	c.unlink(node)
	c.linkHead(node)
}

func (c *LRUCache) linkHead(node *cacheNode) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *LRUCache) unlink(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
