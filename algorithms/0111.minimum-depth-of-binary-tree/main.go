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
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
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
