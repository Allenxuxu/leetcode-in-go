package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(root *TreeNode, min, max *int) bool {
	if root != nil {
		if min != nil && root.Val <= *min {
			return false
		}
		if max != nil && root.Val >= *max {
			return false
		}

		if !helper(root.Right, &root.Val, max) {
			return false
		}
		if !helper(root.Left, min, &root.Val) {
			return false
		}
	}

	return true
}
