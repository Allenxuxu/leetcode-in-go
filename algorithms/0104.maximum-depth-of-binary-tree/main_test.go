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
		[]int{},
		[]int{},
		0,
	},

	{
		[]int{3, 9, 20, 15, 7},
		[]int{9, 3, 15, 20, 7},
		3,
	},

	{
		[]int{3, 9, 20, 15, 10, 7},
		[]int{9, 3, 10, 15, 20, 7},
		4,
	},

	// 可以多个 testcase
}

func Test_maxDepth(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxDepth(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func Test_maxDepth1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxDepth1(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func Test_maxDepth2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, maxDepth2(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
