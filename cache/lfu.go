package cache

import "container/list"

type lfuCacheNode struct {
	key, value, freq int
}

type LFUCache struct {
	capacity, minFreq int
	key2Node          map[int]*list.Element
	freq2List         map[int]*list.List
}

func NewLFUCache(capacity int) LFUCache {
	return LFUCache{
		capacity:  capacity,
		key2Node:  map[int]*list.Element{},
		freq2List: map[int]*list.List{},
	}
}

func (c *LFUCache) pushFront(node *lfuCacheNode) {
	if _, ok := c.freq2List[node.freq]; !ok {
		c.freq2List[node.freq] = list.New()
	}
	c.key2Node[node.key] = c.freq2List[node.freq].PushFront(node)
}

func (c *LFUCache) getNode(key int) *lfuCacheNode {
	element, ok := c.key2Node[key]
	if !ok {
		return nil
	}

	node := element.Value.(*lfuCacheNode)
	freqList := c.freq2List[node.freq]
	freqList.Remove(element)
	if freqList.Len() == 0 {
		delete(c.freq2List, node.freq)
		if c.minFreq == node.freq {
			c.minFreq++
		}
	}
	node.freq++
	c.pushFront(node)

	return node
}

func (c *LFUCache) Get(key int) (int, bool) {
	if node := c.getNode(key); node != nil {
		return node.value, true
	}

	return -1, false
}

func (c *LFUCache) Put(key, value int) {
	if node := c.getNode(key); node != nil {
		node.value = value
		return
	}

	if len(c.key2Node) == c.capacity {
		leftList := c.freq2List[c.minFreq]
		temp := leftList.Remove(leftList.Back()).(*lfuCacheNode)
		delete(c.key2Node, temp.key)
		if leftList.Len() == 0 {
			delete(c.freq2List, c.minFreq)
		}
	}
	c.pushFront(&lfuCacheNode{key: key, value: value, freq: 1})
	c.minFreq = 1
}
