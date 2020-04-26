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
