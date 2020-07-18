package algorithm

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return dfs(root) != -1
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)
	if left == -1 {
		return -1
	}
	if right == -1 {
		return -1
	}

	t := int(math.Abs(float64(left - right)))
	if t < 2 {
		return max(left, right) + 1
	} else {
		return -1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
