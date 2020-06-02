package algorithm

import "github.com/Allenxuxu/toolkit/stack"

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

func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	s := stack.New()
	s.Push(root)

	for s.Len() > 0 {
		node := s.Pop().(*TreeNode)
		ret = append(ret, node.Val)

		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}

	return ret
}

func preorderTraversal2(root *TreeNode) []int {
	var ret []int
	s := stack.New()

	for root != nil || s.Len() > 0 {
		for root != nil {
			ret = append(ret, root.Val)

			s.Push(root)
			root = root.Left
		}

		root = s.Pop().(*TreeNode)
		root = root.Right
	}

	return ret
}

// 推荐模版
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	s := stack.New()
	s.Push(root)

	for s.Len() > 0 {
		c := s.Pop()

		if c != nil {
			node := c.(*TreeNode)
			if node.Right != nil { //右节点先压栈，最后处理
				s.Push(node.Right)
			}
			if node.Left != nil {
				s.Push(node.Left)
			}

			s.Push(node) //当前节点重新压栈（留着以后处理），因为先序遍历所以最后压栈
			s.Push(nil)  //在当前节点之前加入一个空节点表示已经访问过了
		} else { // 当前 c == nil , 说明这个节点已经访问过了
			node := s.Pop().(*TreeNode) // node 是上面 s.Push(node) 中的那个 node
			ret = append(ret, node.Val)
		}
	}

	return ret
}
