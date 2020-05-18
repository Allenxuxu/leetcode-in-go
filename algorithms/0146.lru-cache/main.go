package algorithm

import "container/list"

type Node struct {
	Key, Value int
}

type LRUCache struct {
	m    map[int]*list.Element
	list *list.List
	cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:    make(map[int]*list.Element, capacity),
		list: list.New(),
		cap:  capacity,
	}
}

func (l *LRUCache) Get(key int) int {
	v, ok := l.m[key]
	if !ok {
		return -1
	}

	l.list.Remove(v)
	node := v.Value.(*Node)
	l.m[key] = l.list.PushFront(node)

	return node.Value
}

func (l *LRUCache) Put(key int, value int) {
	node := Node{
		Key:   key,
		Value: value,
	}

	v, ok := l.m[key]
	if !ok {
		e := l.list.PushFront(&node)
		l.m[key] = e

		if l.list.Len() > l.cap {
			toDel := l.list.Back()
			delete(l.m, toDel.Value.(*Node).Key)
			l.list.Remove(toDel)
		}
	} else {
		l.list.Remove(v)
		e := l.list.PushFront(&node)
		l.m[key] = e
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
