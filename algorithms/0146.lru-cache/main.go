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

func (this *LRUCache) Get(key int) int {
	e, ok := this.data[key]
	if ok {
		this.q.MoveToFront(e)
		return e.Value.(*KeyValue).value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.data[key]
	if !ok {
		if this.q.Len() == this.c {
			b := this.q.Back()
			delete(this.data, b.Value.(*KeyValue).key)
			this.q.Remove(b)
		}
		e = this.q.PushFront(&KeyValue{key: key, value: value})
		this.data[key] = e
	} else {
		e.Value.(*KeyValue).value = value
		this.q.MoveToFront(e)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
