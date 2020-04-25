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

func postOrderTraversal(root *TreeNode) []int {
	var ret []int
	helperPostOrder(root, &ret)
	return ret
}

func helperPostOrder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	*data = append(*data, root.Val)
	helperPostOrder(root.Left, data)
	helperPostOrder(root.Right, data)
}
