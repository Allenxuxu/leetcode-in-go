package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{preorder[0], nil, nil}

	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

//对于任意一颗树而言，前序遍历的形式总是
//
//[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]
//即根节点总是前序遍历中的第一个节点。而中序遍历的形式总是
//
//[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]
//
