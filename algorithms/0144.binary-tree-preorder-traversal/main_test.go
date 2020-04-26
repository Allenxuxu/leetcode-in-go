package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in, post []int
}{

	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{3, 2, 1},
	},

	{
		[]int{1, 2, 4, 5, 3, 6, 7},
		[]int{4, 2, 5, 1, 6, 3, 7},
		[]int{4, 5, 2, 6, 7, 3, 1},
	},

	// 可以有多个 testcase
}

func Test_preorderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := InPost2Tree(tc.in, tc.post)
		ast.Equal(tc.pre, preorderTraversal(root), "输入:%v", tc)
	}
}

func Test_preorderTraversal1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := InPost2Tree(tc.in, tc.post)
		ast.Equal(tc.pre, preorderTraversal1(root), "输入:%v", tc)
	}
}
