package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	pre  []int
}{

	{
		[]int{},
		nil,
	},

	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{4, 2, 1, 3, 6, 5, 7},
	},

	// 可以多个 testcase
}

func Test_Problem0108(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.pre, Tree2Preorder(sortedArrayToBST(tc.nums)), "输入:%v", tc)
	}
}
