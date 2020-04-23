package algorithm

type MyCircularDeque struct {
	array []int

	front int
	rear  int
	cap   int
	len   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		array: make([]int, k),
		front: 0,
		rear:  1,
		cap:   k,
		len:   0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.len < d.cap {
		d.array[d.front] = value
		d.front = (d.front - 1 + d.cap) % d.cap
		d.len++
		return true
	}
	return false
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.len < d.cap {
		d.array[d.rear] = value
		d.rear = (d.rear + 1) % d.cap
		d.len++
		return true
	}
	return false
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) DeleteFront() bool {
	if d.len > 0 {
		d.front = (d.front + 1) % d.cap
		d.len--
		return true
	}
	return false
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) DeleteLast() bool {
	if d.len > 0 {
		d.rear = (d.rear - 1 + d.cap) % d.cap
		d.len--
		return true
	}
	return false
}

/** Get the front item from the deque. */
func (d *MyCircularDeque) GetFront() int {
	if d.len > 0 {
		return d.array[(d.front+1)%d.cap]
	}
	return -1
}

/** Get the last item from the deque. */
func (d *MyCircularDeque) GetRear() int {
	if d.len > 0 {
		return d.array[(d.rear-1+d.cap)%d.cap]
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (d *MyCircularDeque) IsEmpty() bool {
	return d.len == 0
}

/** Checks whether the circular deque is full or not. */
func (d *MyCircularDeque) IsFull() bool {
	return d.len == d.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
