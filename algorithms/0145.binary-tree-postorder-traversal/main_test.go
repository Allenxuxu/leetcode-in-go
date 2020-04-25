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

func Test_postorderTraversal(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		root := PreIn2Tree(tc.pre, tc.in)
		ast.Equal(tc.post, postorderTraversal(root), "输入:%v", tc)
	}
}
