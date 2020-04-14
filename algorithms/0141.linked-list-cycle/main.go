package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var (
		f = head.Next
		s = head
	)

	for s != f {
		if s == nil || f == nil || f.Next == nil {
			return false
		}

		s = s.Next
		f = f.Next.Next
	}
	return true
}

func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]bool)

	for head != nil {
		if m[head] {
			return true
		}

		m[head] = true
		head = head.Next
	}

	return false
}
