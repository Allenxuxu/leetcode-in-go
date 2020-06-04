package algorithm

import (
	"github.com/Allenxuxu/toolkit/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type Entry struct {
		*TreeNode
		Level int
	}

	queue := queue.New()
	queue.Push(&Entry{
		TreeNode: root,
		Level:    1,
	})
	var ret int
	for queue.Len() != 0 {
		node := queue.Pop().(*Entry)
		ret = node.Level
		if node.Left != nil {
			queue.Push(&Entry{
				TreeNode: node.Left,
				Level:    node.Level + 1,
			})
		}
		if node.Right != nil {
			queue.Push(&Entry{
				TreeNode: node.Right,
				Level:    node.Level + 1,
			})
		}
		if node.Left == nil && node.Right == nil {
			break
		}
	}

	return ret
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ret int
	q := queue.New()
	q.Push(root)

	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			node := q.Pop().(*TreeNode)

			if node.Left == nil && node.Right == nil {
				return ret + 1
			}
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}

		ret++
	}
	return ret
}
