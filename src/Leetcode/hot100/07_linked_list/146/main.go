package main

import "container/list"

type entry struct {
	key, val int
}

type LRUCache struct {
	capacity  int
	list      *list.List // 双向链表
	keyToNode map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, list: list.New(), keyToNode: make(map[int]*list.Element)}
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	c.list.MoveToFront(node) // 放在最前面
	return node.Value.(entry).val
}

func (c *LRUCache) Put(key int, value int) {
	if node := c.keyToNode[key]; node != nil { // 存在更新
		node.Value = entry{key, value}
		c.list.MoveToFront(node)
		return
	}
	// 不存在添加
	c.keyToNode[key] = c.list.PushFront(entry{key, value})
	if len(c.keyToNode) > c.capacity { // 爆了 删掉
		delete(c.keyToNode, c.list.Remove(c.list.Back()).(entry).key)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
