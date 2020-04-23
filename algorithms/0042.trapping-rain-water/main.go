package algorithm

import (
	"container/list"
)

type stack struct {
	list *list.List
}

func New() *stack {
	return &stack{
		list: list.New(),
	}
}

func (q *stack) Push(i int) {
	q.list.PushBack(i)
}

func (q *stack) Peek() int {
	return q.list.Back().Value.(int)
}
func (q *stack) Pop() int {
	e := q.list.Back()
	q.list.Remove(e)
	return e.Value.(int)
}

func (q *stack) Len() int {
	return q.list.Len()
}

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	q := New()
	var ret int
	for i := 0; i < len(height); i++ {
		for q.Len() > 0 && height[i] > height[q.Peek()] {
			minH := height[q.Pop()]
			if q.Len() == 0 {
				break
			}
			left, right := q.Peek(), i
			ret += (right - left - 1) * (min(height[left], height[right]) - minH)
		}

		q.Push(i)
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap1(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	var (
		ret      int
		leftMax  int
		rightMax int
	)
	for left, right := 0, len(height)-1; left < right; {
		if height[left] < height[right] {
			ret += max(0, leftMax-height[left])
			leftMax = max(height[left], leftMax)
			left++
		} else {
			ret += max(0, rightMax-height[right])
			rightMax = max(height[right], rightMax)
			right--
		}
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
