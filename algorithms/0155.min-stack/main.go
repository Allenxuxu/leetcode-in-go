package algorithm

import "container/list"

type listEntry struct {
	Min int
	X   int
}
type MinStack struct {
	list *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		list: list.New(),
	}
}

func (s *MinStack) Push(x int) {
	var min = x
	if s.list.Len() > 0 {
		lastMin := s.list.Back().Value.(listEntry).Min
		if x > lastMin {
			min = lastMin
		}
	}

	s.list.PushBack(listEntry{
		Min: min,
		X:   x,
	})
}

func (s *MinStack) Pop() {
	if s.list.Len() > 0 {
		s.list.Remove(s.list.Back())
	}
}

func (s *MinStack) Top() int {
	if s.list.Len() > 0 {
		return s.list.Back().Value.(listEntry).X
	}
	return 0
}

func (s *MinStack) GetMin() int {
	if s.list.Len() > 0 {
		return s.list.Back().Value.(listEntry).Min
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
