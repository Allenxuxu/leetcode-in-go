package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{

	{
		499979,
		41537,
	},

	{
		2,
		0,
	},

	{
		5,
		2,
	},

	{
		10,
		4,
	},

	{
		100,
		25,
	},

	// 可以有多个 testcase
}

func Test_countPrimes(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countPrimes(tc.n), "输入:%v", tc)
	}
}

func Test_countPrimes1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, countPrimes1(tc.n), "输入:%v", tc)
	}
}
