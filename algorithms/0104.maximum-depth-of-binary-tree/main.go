package algorithm

import (
	"github.com/Allenxuxu/toolkit/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	var ret int
	if left > right {
		ret = left
	} else {
		ret = right
	}

	return ret + 1
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type Entry struct {
		*TreeNode
		Level int
	}

	var (
		ret int
	)
	queue := queue.New()
	queue.Push(&Entry{
		TreeNode: root,
		Level:    1,
	})
	for queue.Len() != 0 {
		node := queue.Pop().(*Entry)
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

		if node.Level > ret {
			ret = node.Level
		}
	}

	return ret
}
