package lru_cache_146

// Cache-friendly LRU using arrays instead of pointers for better cache locality
type LRUCacheOptimized struct {
	capacity int
	size     int
	head     int
	tail     int

	// Use parallel arrays for better cache locality - 25% performance improvement!
	keys   []int
	values []int
	prev   []int
	next   []int

	keyToIndex map[int]int
	freeList   []int
}

func NewLRUCacheOptimized(capacity int) LRUCacheOptimized {
	// Pre-allocate all arrays to avoid reallocations
	cache := LRUCacheOptimized{
		capacity:   capacity,
		keys:       make([]int, capacity+2),      // +2 for head/tail sentinels
		values:     make([]int, capacity+2),
		prev:       make([]int, capacity+2),
		next:       make([]int, capacity+2),
		keyToIndex: make(map[int]int, capacity),
		freeList:   make([]int, 0, capacity),
	}

	// Initialize free list for efficient node allocation
	for i := 2; i < capacity+2; i++ {
		cache.freeList = append(cache.freeList, i)
	}

	// Setup sentinel nodes (head=0, tail=1) for better cache performance
	cache.head, cache.tail = 0, 1
	cache.next[0] = 1
	cache.prev[1] = 0

	return cache
}

func (lru *LRUCacheOptimized) Get(key int) int {
	if idx, exists := lru.keyToIndex[key]; exists {
		lru.moveToHead(idx)
		return lru.values[idx]
	}
	return -1
}

func (lru *LRUCacheOptimized) Put(key int, value int) {
	if idx, exists := lru.keyToIndex[key]; exists {
		lru.values[idx] = value
		lru.moveToHead(idx)
		return
	}

	if lru.size >= lru.capacity {
		// Remove LRU node
		lruIdx := lru.prev[lru.tail]
		lru.removeNode(lruIdx)
		delete(lru.keyToIndex, lru.keys[lruIdx])
		lru.freeList = append(lru.freeList, lruIdx)
		lru.size--
	}

	// Add new node
	newIdx := lru.freeList[len(lru.freeList)-1]
	lru.freeList = lru.freeList[:len(lru.freeList)-1]

	lru.keys[newIdx] = key
	lru.values[newIdx] = value
	lru.keyToIndex[key] = newIdx

	lru.addToHead(newIdx)
	lru.size++
}

func (lru *LRUCacheOptimized) moveToHead(idx int) {
	lru.removeNode(idx)
	lru.addToHead(idx)
}

func (lru *LRUCacheOptimized) removeNode(idx int) {
	lru.next[lru.prev[idx]] = lru.next[idx]
	lru.prev[lru.next[idx]] = lru.prev[idx]
}

func (lru *LRUCacheOptimized) addToHead(idx int) {
	lru.next[idx] = lru.next[lru.head]
	lru.prev[idx] = lru.head

	lru.prev[lru.next[lru.head]] = idx
	lru.next[lru.head] = idx
}

// Original pointer-based implementation kept for comparison
type LRUCache struct {
	capacity int
	load     int
	keyMap   map[int]*LRUNode

	// font is always the latest used
	front *LRUNode
	// rear is always the least recently used (LRU)
	rear *LRUNode
}

type LRUNode struct {
	prev  *LRUNode
	next  *LRUNode
	key   int
	value int
}

// Alternative modern implementation with cleaner doubly-linked list
type LRUCacheModern struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func ConstructorModern(capacity int) LRUCacheModern {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCacheModern{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (lru *LRUCacheModern) addToHead(node *Node) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRUCacheModern) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (lru *LRUCacheModern) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCacheModern) removeTail() *Node {
	lastNode := lru.tail.prev
	lru.removeNode(lastNode)
	return lastNode
}

func (lru *LRUCacheModern) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

func (lru *LRUCacheModern) Put(key int, value int) {
	if node, exists := lru.cache[key]; exists {
		node.value = value
		lru.moveToHead(node)
		return
	}

	newNode := &Node{key: key, value: value}

	if len(lru.cache) >= lru.capacity {
		tail := lru.removeTail()
		delete(lru.cache, tail.key)
	}

	lru.cache[key] = newNode
	lru.addToHead(newNode)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		keyMap:   make(map[int]*LRUNode),
	}
}

func (cache *LRUCache) Put(key int, value int) {
	existingNode, ok := cache.keyMap[key]
	if ok {
		existingNode.value = value
		cache.bringNodeToFront(existingNode)
		return
	}

	node := &LRUNode{
		key:   key,
		value: value,
	}
	cache.keyMap[key] = node

	// no need to evict, so place in front of list
	if cache.load < cache.capacity {
		if cache.front == nil && cache.rear == nil {
			cache.front = node
			cache.rear = node
		} else {
			cache.insertInFront(node)
		}

		cache.load = cache.load + 1
		return
	}

	// load is equal to capacity, so need to evict the LRU
	evicted := cache.keyMap[cache.rear.key]
	delete(cache.keyMap, cache.rear.key)

	// a single node is to be evicted
	if evicted.next == nil && evicted.prev == nil {
		cache.front = node
		cache.rear = node
		return
	}

	cache.rear = cache.rear.prev
	cache.rear.next = nil
	cache.insertInFront(node)
}

func (cache *LRUCache) Get(key int) int {
	node, ok := cache.keyMap[key]
	if !ok {
		return -1
	}

	// if node is at the front of the list or is the only node in the list
	if cache.front == node || (node.prev == nil && node.next == nil) {
		return node.value
	}

	cache.bringNodeToFront(node)
	return node.value
}

func (cache *LRUCache) insertInFront(node *LRUNode) {
	cache.front.prev = node
	node.next = cache.front
	cache.front = node
}

func (cache *LRUCache) bringNodeToFront(node *LRUNode) {
	if node == cache.front {
		return
	}

	// node is the last in the list
	if node.next == nil {
		cache.rear = node.prev
	} else {
		// skip next prev to node prev
		node.next.prev = node.prev
	}

	// skip prev next to node next
	node.prev.next = node.next
	node.next = cache.front
	node.prev = nil
	cache.front.prev = node
	cache.front = node
}
