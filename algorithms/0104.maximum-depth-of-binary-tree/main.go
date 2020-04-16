package algorithm

import "container/list"

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
	queue := list.New()
	queue.PushBack(&Entry{
		TreeNode: root,
		Level:    1,
	})
	for queue.Len() != 0 {
		item := queue.Front()
		node := item.Value.(*Entry)
		queue.Remove(item)
		if node.Left != nil {
			queue.PushBack(&Entry{
				TreeNode: node.Left,
				Level:    node.Level + 1,
			})
		}
		if node.Right != nil {
			queue.PushBack(&Entry{
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
