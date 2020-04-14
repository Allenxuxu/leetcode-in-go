package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	ints []int
	pos  int
	ans  bool
}{

	{
		[]int{-21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5},
		-1,
		false,
	},

	{
		[]int{3, 2, 0, -4},
		1,
		true,
	},

	{
		[]int{1, 2},
		0,
		true,
	},

	{
		[]int{},
		-1,
		false,
	},

	{
		[]int{1, 2},
		-1,
		false,
	},

	{
		[]int{1},
		-1,
		false,
	},

	// 可以有多个 testcase
}

func Test_hasCycle(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		head := Ints2ListWithCycle(tc.ints, tc.pos)
		ast.Equal(tc.ans, hasCycle(head), "输入:%v", tc)
	}
}

func Test_hasCycle1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		head := Ints2ListWithCycle(tc.ints, tc.pos)
		ast.Equal(tc.ans, hasCycle1(head), "输入:%v", tc)
	}
}
