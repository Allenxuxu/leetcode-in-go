package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var dummy = new(ListNode)
	dummy.Next = head
	pre, first, second := dummy, head, head.Next
	for first != nil && second != nil {
		pre.Next = second
		second.Next, first.Next = first, second.Next
		pre, first = pre.Next.Next, pre.Next.Next.Next
		if first != nil {
			second = first.Next
		}
	}
	return dummy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs1(next.Next)
	next.Next = head

	return next
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
