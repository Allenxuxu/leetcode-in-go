package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  bool
}{

	{[]int{1, 2, 3, 2, 1}, true},
	{[]int{1, 3, 2, 1}, false},

	// 可以有多个 testcase
}

func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := Ints2List(tc.nums)
		ast.Equal(tc.ans, isPalindrome(head), "输入:%v", tc)
	}
}

func Test_isPalindrome1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := Ints2List(tc.nums)
		ast.Equal(tc.ans, isPalindrome1(head), "输入:%v", tc)
	}
}

func Test_isPalindrome2(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		head := Ints2List(tc.nums)
		ast.Equal(tc.ans, isPalindrome2(head), "输入:%v", tc)
	}
}
