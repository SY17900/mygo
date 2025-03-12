package cache

import "container/list"

type lruCacheNode struct {
	key, value int
}

type LRUCache struct {
	capacity int
	mp       map[int]*list.Element
	lst      *list.List
}

func NewLRUCache(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		mp:       make(map[int]*list.Element),
		lst:      list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if e, ok := c.mp[key]; ok {
		c.lst.MoveToFront(e)
		node := e.Value.(*lruCacheNode)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if e, ok := c.mp[key]; ok {
		node := e.Value.(*lruCacheNode)
		node.value = value
		c.lst.MoveToFront(e)
	} else {
		if c.lst.Len() == c.capacity {
			back := c.lst.Back()
			if back != nil {
				temp := c.lst.Remove(back).(*lruCacheNode)
				delete(c.mp, temp.key)
			}
		}
		e := c.lst.PushFront(&lruCacheNode{key: key, value: value})
		c.mp[key] = e
	}
}
