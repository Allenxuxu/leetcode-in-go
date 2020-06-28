package algorithm

import "github.com/Allenxuxu/toolkit/queue"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	q := queue.New()
	q.Push(root)

	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			node := q.Pop().(*TreeNode)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}

			if i == size-1 {
				ret = append(ret, node.Val)
			}
		}
	}

	return ret
}

func rightSideView1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	dfs(root, 0, &ret)
	return ret
}

func dfs(root *TreeNode, depth int, ret *[]int) {
	if root != nil {
		if depth == len(*ret) {
			*ret = append(*ret, root.Val)
		}

		depth++
		dfs(root.Right, depth, ret)
		dfs(root.Left, depth, ret)
	}
}
