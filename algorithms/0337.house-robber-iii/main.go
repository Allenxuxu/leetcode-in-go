package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	ret := dfs(root)
	return max(ret[0], ret[1])
}

func dfs(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{}
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	// dp[0]：以当前 node 为根结点的子树能够偷取的最大价值，规定 node 结点不偷
	// dp[1]：以当前 node 为根结点的子树能够偷取的最大价值，规定 node 结点偷
	var dp [2]int

	dp[0] = max(left[0], left[1]) + max(right[0], right[1])
	dp[1] = root.Val + left[0] + right[0]

	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
