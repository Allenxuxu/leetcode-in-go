package algorithm

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
