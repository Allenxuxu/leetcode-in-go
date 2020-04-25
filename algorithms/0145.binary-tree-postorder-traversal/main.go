package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var ret []int
	helperPostOrder(root, &ret)
	return ret
}

func helperPostOrder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	helperPostOrder(root.Left, data)
	helperPostOrder(root.Right, data)
	*data = append(*data, root.Val)
}
