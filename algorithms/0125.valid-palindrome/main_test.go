package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans bool
}{

	{
		"A man, a plan, a canal: Panama",
		true,
	},

	{
		"race a car",
		false,
	},
	{
		"0P",
		false,
	},
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isPalindrome(tc.s), "输入:%v", tc)
	}
}
