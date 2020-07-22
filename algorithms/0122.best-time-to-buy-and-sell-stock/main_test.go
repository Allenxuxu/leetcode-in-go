package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	nums []int
	ans  int
}{
	{
		[]int{},
		0,
	},
	{
		[]int{7, 1, 5, 3, 6, 4},
		7,
	},
	{
		[]int{7, 6, 4, 3, 1},
		0,
	},
	{
		[]int{1, 2, 3, 4, 5},
		4,
	},
	{
		[]int{1, 2, 10, 11, 12, 15},
		14,
	},
}

func Test_maxProfit(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxProfit(tc.nums), "输入:%v", tc)
	}
}

func Test_maxProfit1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, maxProfit1(tc.nums), "输入:%v", tc)
	}
}
