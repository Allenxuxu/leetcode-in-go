package algorithm

import "github.com/Allenxuxu/toolkit/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	helperPreOrder(root, &ret)
	return ret
}

func helperPreOrder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	*data = append(*data, root.Val)
	helperPreOrder(root.Left, data)
	helperPreOrder(root.Right, data)
}

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	s := stack.New()
	s.Push(root)

	for s.Len() > 0 {
		node := s.Pop().(*TreeNode)
		ret = append(ret, node.Val)

		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}

	return ret
}

func preorderTraversal2(root *TreeNode) []int {
	var ret []int
	s := stack.New()

	for root != nil || s.Len() > 0 {
		for root != nil {
			s.Push(root)
			ret = append(ret, root.Val)
			root = root.Left
		}

		root = s.Pop().(*TreeNode)
		root = root.Right
	}

	return ret
}
