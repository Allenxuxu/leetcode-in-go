package algorithm

import "container/list"

type deque struct {
	list *list.List
}

func New() *deque {
	return &deque{
		list: list.New(),
	}
}

func (q *deque) PushBack(i int) {
	q.list.PushBack(i)
}

func (q *deque) PushFront(i int) {
	q.list.PushFront(i)
}

func (q *deque) Front() int {
	return q.list.Front().Value.(int)
}

func (q *deque) Back() int {
	return q.list.Back().Value.(int)
}

func (q *deque) PopFront() int {
	e := q.list.Front()
	q.list.Remove(e)
	return e.Value.(int)
}

func (q *deque) PopBack() int {
	e := q.list.Back()
	q.list.Remove(e)
	return e.Value.(int)
}

func (q *deque) Len() int {
	return q.list.Len()
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ret := make([]int, 0, len(nums)-k)
	q := New()
	for i := 0; i < k; i++ {
		for q.Len() != 0 && nums[q.Back()] < nums[i] {
			q.PopBack()
		}
		q.PushBack(i)
	}
	ret = append(ret, nums[q.Front()])

	for i := k; i < len(nums); i++ {
		if nums[q.Front()] == nums[i-k] {
			q.PopFront()
		}

		for q.Len() != 0 && nums[q.Back()] < nums[i] {
			q.PopBack()
		}
		q.PushBack(i)

		ret = append(ret, nums[q.Front()])
	}
	return ret
}
