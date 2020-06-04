package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	pre, in []int
	ans     int
}{

	{
		[]int{1, 2, 3, 4, 5},
		[]int{5, 4, 3, 2, 1},
		5,
	},

	{
		[]int{1, 2},
		[]int{2, 1},
		2,
	},

	{
		[]int{1},
		[]int{1},
		1,
	},

	{
		[]int{},
		[]int{},
		0,
	},

	{
		[]int{2, 1, 5, 4, 3, 6, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		2,
	},

	{
		[]int{4, 2, 1, 3, 6, 5, 7},
		[]int{1, 2, 3, 4, 5, 6, 7},
		3,
	},

	// 可以多个 testcase
}

func Test_minDepth(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minDepth(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func Test_minDepth1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minDepth1(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func Test_minDepth2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, minDepth2(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
