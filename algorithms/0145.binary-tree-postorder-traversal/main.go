package algorithm

import (
	"github.com/Allenxuxu/toolkit/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ret []int
	helperPostOrder(root, &ret)
	return ret
}

func helperPostOrder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	helperPostOrder(root.Left, data)
	helperPostOrder(root.Right, data)
	*data = append(*data, root.Val)
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	s := stack.New()
	s.Push(root)

	for s.Len() > 0 {
		node := s.Pop().(*TreeNode)
		ret = append(ret, node.Val)

		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
	}

	reversal(ret)
	return ret
}

func reversal(a []int) {
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
}

func postorderTraversal2(root *TreeNode) []int {
	var ret []int
	s := stack.New()
	var last *TreeNode
	for root != nil || s.Len() > 0 {
		for root != nil {
			s.Push(root)
			root = root.Left
		}

		// peek
		root = s.Pop().(*TreeNode)
		s.Push(root)

		if root.Right == nil || root.Right == last {
			ret = append(ret, root.Val)
			s.Pop()

			last = root
			root = nil
		} else {
			root = root.Right
		}
	}
	return ret
}

// 推荐模版
func postorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	s := stack.New()
	s.Push(root)

	for s.Len() > 0 {
		c := s.Pop()

		if c != nil {
			node := c.(*TreeNode)

			s.Push(node)
			s.Push(nil)

			if node.Right != nil {
				s.Push(node.Right)
			}
			if node.Left != nil {
				s.Push(node.Left)
			}
		} else {
			node := s.Pop().(*TreeNode)
			ret = append(ret, node.Val)
		}
	}

	return ret
}
