package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head
	pre, end := dummy, dummy
	for {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		start, next := pre.Next, end.Next
		pre.Next = reverseList(start, next)
		start.Next = next
		pre, end = start, start
	}
	return dummy.Next
}

func reverseList(head *ListNode, end *ListNode) *ListNode {
	var prev *ListNode
	for head != end {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
