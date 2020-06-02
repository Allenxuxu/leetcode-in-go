package algorithm

import "github.com/Allenxuxu/toolkit/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	helper(root, &ret)
	return ret
}

func helper(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	helper(root.Left, data)
	*data = append(*data, root.Val)
	helper(root.Right, data)
}

func inorderTraversal2(root *TreeNode) []int {
	var ret []int
	s := stack.New()

	for root != nil || s.Len() > 0 {
		for root != nil {
			s.Push(root)
			root = root.Left
		}

		root = s.Pop().(*TreeNode)
		ret = append(ret, root.Val)

		root = root.Right
	}

	return ret
}

// 推荐模版
func inorderTraversal3(root *TreeNode) []int {
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

			if node.Right != nil {
				s.Push(node.Right) //右节点先压栈，最后处理
			}
			s.Push(node)
			s.Push(nil)
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
