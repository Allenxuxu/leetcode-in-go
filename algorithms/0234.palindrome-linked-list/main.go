package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

//  将元素复制到数组，然后双指针判断数组是不是回文数组
func isPalindrome(head *ListNode) bool {
	var tmp []int

	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}

	for i, j := 0, len(tmp)-1; i < j; {
		if tmp[i] != tmp[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// 快慢指针，慢指针步进 1， 快指针步进 2，翻转链表前半部分，然后对比
func isPalindrome1(head *ListNode) bool {
	var (
		f = head
		s = head
	)

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	p := reverseList(s)
	for head != nil && p != nil {
		if p.Val != head.Val {
			return false
		}
		p = p.Next
		head = head.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// 递归
var _current *ListNode

func isPalindrome2(head *ListNode) bool {
	_current = head

	return helper(head)
}

func helper(p *ListNode) bool {
	if p == nil {
		return true
	}

	if !helper(p.Next) {
		return false
	}

	if p.Val != _current.Val {
		return false
	}

	_current = _current.Next
	return true
}
