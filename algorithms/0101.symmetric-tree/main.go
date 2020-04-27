package algorithm

import "github.com/Allenxuxu/toolkit/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return (left.Val == right.Val) && isMirror(left.Left, right.Right) && isMirror(right.Left, left.Right)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	s := stack.New()
	s.Push(root.Left)
	s.Push(root.Right)

	for s.Len() > 0 {
		var (
			left  *TreeNode
			right *TreeNode
		)
		if node := s.Pop(); node != nil {
			left = node.(*TreeNode)
		}
		if node := s.Pop(); node != nil {
			right = node.(*TreeNode)
		}

		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		s.Push(left.Left)
		s.Push(right.Right)
		s.Push(left.Right)
		s.Push(right.Left)
	}

	return true
}
