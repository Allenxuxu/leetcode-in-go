package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	a, b   []int
	ea, eb int
}{

	{
		[]int{4, 1, 8, 4, 5},
		[]int{5, 0, 1, 8, 4, 5},
		2,
		3,
	},

	{
		[]int{0, 9, 1, 2, 4},
		[]int{3, 2, 4},
		3,
		1,
	},

	{
		[]int{2, 6, 4},
		[]int{1, 5},
		3,
		2,
	},

	// 可以有多个 testcase
}

// head must Not be nil
func tailOf(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func Test_getIntersectionNode(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		tail := Ints2List(tc.a[tc.ea:])

		headA := Ints2List(tc.a[:tc.ea])
		tailA := tailOf(headA)
		tailA.Next = tail

		headB := Ints2List(tc.b[:tc.eb])
		tailB := tailOf(headB)
		tailB.Next = tail

		ast.Equal(tail, getIntersectionNode(headA, headB), "输入:%v", tc)
	}
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
