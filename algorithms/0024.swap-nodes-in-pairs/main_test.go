package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

// para 是参数
// one 代表第一个参数
type para struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one []int
}

var qs = []question{

	question{
		para{[]int{}},
		ans{[]int{}},
	},

	question{
		para{[]int{1}},
		ans{[]int{1}},
	},

	question{
		para{[]int{1, 2, 3, 4}},
		ans{[]int{2, 1, 4, 3}},
	},

	question{
		para{[]int{1, 2, 3, 4, 5}},
		ans{[]int{2, 1, 4, 3, 5}},
	},

	// 如需多个测试，可以复制上方元素。
}

func Test_swapPairs(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(swapPairs(s2l(p.one))), "输入:%v", p)
	}
}

func Test_swapPairs1(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(swapPairs1(s2l(p.one))), "输入:%v", p)
	}
}

func Test_swapPairs2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, l2s(swapPairs2(s2l(p.one))), "输入:%v", p)
	}
}

// convert *ListNode to []int
func l2s(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// convert []int to *ListNode
func s2l(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res
	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{
			Val: nums[i],
		}
		temp = temp.Next
	}

	return res
}
