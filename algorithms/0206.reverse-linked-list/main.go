package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	var pre *ListNode
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre
}
