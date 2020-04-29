package algorithm

import (
	"github.com/Allenxuxu/toolkit/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	q := queue.New()
	q.Push(root)

	for q.Len() > 0 {
		node := q.Pop().(*TreeNode)

		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}
	return root
}
