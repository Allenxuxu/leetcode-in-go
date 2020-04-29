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

func isValidBST1(root *TreeNode) bool {
	var ret []int
	inorder(root, &ret)

	for i := 1; i < len(ret); i++ {
		if ret[i-1] >= ret[i] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	inorder(root.Left, ret)
	*ret = append(*ret, root.Val)
	inorder(root.Right, ret)
}
