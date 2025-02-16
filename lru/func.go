package lru

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
		node := NewcacheNode(key, value)
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

func (cache *LRUCache) insertNew(node *cacheNode) {
	cache.insertHead(node)
	cache.size++
}

func (cache *LRUCache) updateOld(node *cacheNode) {
	cache.remove(node)
	cache.insertHead(node)
}

func (cache *LRUCache) remove(node *cacheNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (cache *LRUCache) insertHead(node *cacheNode) {
	node.next = cache.head.next
	node.next.pre = node
	cache.head.next = node
	node.pre = cache.head
}
