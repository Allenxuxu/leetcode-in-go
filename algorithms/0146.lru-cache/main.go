package algorithm

import "container/list"

type LRUCache struct {
	q    *list.List
	data map[int]*list.Element
	c    int
}

type KeyValue struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		q:    list.New(),
		data: make(map[int]*list.Element),
		c:    capacity,
	}
	return l
}

func (c *LRUCache) Get(key int) int {
	e, ok := c.data[key]
	if ok {
		// 将节点移动到头部，并且删除
		c.q.MoveToFront(e)
		return e.Value.(*KeyValue).value
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	e, ok := c.data[key]
	if !ok {
		// cache 满了，删除最后一个（最长时间没有使用）
		if c.q.Len() == c.c {
			b := c.q.Back()
			delete(c.data, b.Value.(*KeyValue).key)
			c.q.Remove(b)
		}
		// 插入新节点到头部
		e = c.q.PushFront(&KeyValue{key: key, value: value})
		c.data[key] = e
	} else {
		// 设置节点值
		e.Value.(*KeyValue).value = value
		// 将节点移动到头部
		c.q.MoveToFront(e)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
