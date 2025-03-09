package cache

type lruCacheNode struct {
	key, value int
	pre, next  *lruCacheNode
}

func NewCacheNode(key, value int) *lruCacheNode {
	return &lruCacheNode{key: key, value: value, pre: nil, next: nil}
}

type LRUCache struct {
	size, capacity int
	mp             map[int]*lruCacheNode
	head           *lruCacheNode
}

func NewLRUCache(capacity int) LRUCache {
	cache := LRUCache{
		0,
		capacity,
		map[int]*lruCacheNode{},
		NewCacheNode(0, 0),
	}
	cache.head.next = cache.head
	cache.head.pre = cache.head

	return cache
}

func (cache *LRUCache) Get(key int) int {
	if _, ok := cache.mp[key]; !ok {
		return -1
	}
	node := cache.mp[key]
	cache.updateOld(node)
	return node.value
}

func (cache *LRUCache) Put(key, value int) {
	if _, ok := cache.mp[key]; !ok {
		node := NewCacheNode(key, value)
		for cache.size >= cache.capacity {
			cache.removeLast()
		}
		cache.mp[key] = node
		cache.insertNew(node)
	} else {
		node := cache.mp[key]
		node.value = value
		cache.updateOld(node)
	}
}

func (cache *LRUCache) removeLast() {
	delete(cache.mp, cache.head.pre.key)
	cache.remove(cache.head.pre)
	cache.size--
}

func (cache *LRUCache) insertNew(node *lruCacheNode) {
	cache.insertHead(node)
	cache.size++
}

func (cache *LRUCache) updateOld(node *lruCacheNode) {
	cache.remove(node)
	cache.insertHead(node)
}

func (cache *LRUCache) remove(node *lruCacheNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (cache *LRUCache) insertHead(node *lruCacheNode) {
	node.next = cache.head.next
	node.next.pre = node
	cache.head.next = node
	node.pre = cache.head
}
