package alorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = new(ListNode)
	dummy.Next = head
	end := dummy
	first := dummy
	for i := 0; end != nil; i++ {
		if i > n {
			first = first.Next
		}
		end = end.Next
	}

	toDel := first.Next
	first.Next = toDel.Next
	if toDel.Next != nil {
		toDel.Next = nil
	}

	return dummy.Next
}
