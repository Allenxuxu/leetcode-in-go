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
		[]int{},
		[]int{},
		true,
	},

	{
		[]int{1, 2, 2},
		[]int{2, 1, 2},
		true,
	},

	{
		[]int{1, 2, 3, 2, 3},
		[]int{2, 3, 1, 2, 3},
		false,
	},
}

func Test_Problem0101(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isSymmetric(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}

func Test_Problem0101_1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, isSymmetric1(PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
