package lru

type cacheNode struct {
	key, value int
	pre, next  *cacheNode
}

func NewcacheNode(key, value int) *cacheNode {
	return &cacheNode{key: key, value: value, pre: nil, next: nil}
}

type LRUCache struct {
	size, capacity int
	mp             map[int]*cacheNode
	head           *cacheNode
}

func NewLRUCache(capacity int) *LRUCache {
	cache := LRUCache{
		0,
		capacity,
		map[int]*cacheNode{},
		NewcacheNode(0, 0),
	}
	cache.head.next = cache.head
	cache.head.pre = cache.head

	return &cache
}
