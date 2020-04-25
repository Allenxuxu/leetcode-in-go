package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var ret []int
	helperPreOrder(root, &ret)
	return ret
}

func helperPreOrder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	*data = append(*data, root.Val)
	helperPreOrder(root.Left, data)
	helperPreOrder(root.Right, data)
}
