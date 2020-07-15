package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	preorder []int
	inorder  []int
}

// ans 是答案
type ans struct {
	postorder []int
}

func Test_Problem0105(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		question{
			para{
				[]int{1, 2, 4, 5, 3, 6, 7},
				[]int{4, 2, 5, 1, 6, 3, 7},
			},
			ans{
				[]int{4, 5, 2, 6, 7, 3, 1},
			},
		},

		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.postorder, Tree2Postorder(buildTree(p.preorder, p.inorder)), "输入:%v", p)
	}

	ast.Nil(buildTree([]int{}, []int{}))

	ast.Equal(0, indexOf(3, []int{0, 1, 2}))
}

// Tree2Postorder 把 二叉树 转换成 postorder 的切片
func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}
