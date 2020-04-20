package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 辅助结点指向 head
	var dummy = new(ListNode)
	dummy.Next = head
	pre := dummy
	for head != nil && head.Next != nil {
		second := head.Next

		pre.Next = second
		head.Next = second.Next
		second.Next = head

		pre = head
		head = head.Next
	}

	return dummy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	f := head
	s := head.Next
	f.Next = swapPairs1(s.Next)
	s.Next = f

	return s
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var dummy = new(ListNode)
	dummy.Next = head
	pre, first, second := dummy, head, head.Next
	for first != nil && second != nil {
		pre.Next = second
		second.Next, first.Next, pre, first = first, second.Next, first, first.Next.Next

		if first != nil {
			second = first.Next
		}
	}
	return dummy.Next
}
