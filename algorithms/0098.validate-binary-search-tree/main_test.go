package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     bool
}{

	{
		[]int{1, 2, 3},
		[]int{2, 1, 3},
		false,
	},

	{
		[]int{},
		[]int{},
		true,
	},

	{
		[]int{2, 1, 3},
		[]int{1, 2, 3},
		true,
	},

	{
		[]int{10, 5, 15, 6, 20},
		[]int{5, 10, 6, 15, 20},
		false,
	},

	{
		[]int{3, 2, 1},
		[]int{2, 3, 1},
		false,
	},
	// 可以多个 testcase
}

func Test_isValidBST(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isValidBST(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func Test_isValidBST1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isValidBST1(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
