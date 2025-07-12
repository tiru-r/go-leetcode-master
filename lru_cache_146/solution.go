package lru_cache_146

/* ---------- exported LeetCode interface ---------- */

type LRUCache struct {
	cap  int
	data map[int]*node
	head *node // MRU sentinel
	tail *node // LRU sentinel
}

type node struct {
	prev, next *node
	key, val   int
}

func Constructor(capacity int) LRUCache {
	head, tail := &node{}, &node{} // sentinels
	head.next, tail.prev = tail, head

	return LRUCache{
		cap:  capacity,
		data: make(map[int]*node, capacity),
		head: head,
		tail: tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.data[key]; ok {
		c.moveToHead(n)
		return n.val
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if n, ok := c.data[key]; ok {
		n.val = value
		c.moveToHead(n)
		return
	}

	if len(c.data) == c.cap {
		lru := c.tail.prev
		delete(c.data, lru.key)
		c.remove(lru)
	}

	n := &node{key: key, val: value}
	c.data[key] = n
	c.addToHead(n)
}

/* ---------- internal helpers ---------- */

func (c *LRUCache) moveToHead(n *node) {
	c.remove(n)
	c.addToHead(n)
}

func (c *LRUCache) addToHead(n *node) {
	n.prev = c.head
	n.next = c.head.next
	c.head.next.prev = n
	c.head.next = n
}

func (c *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}
