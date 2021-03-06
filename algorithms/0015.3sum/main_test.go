package algorithm

import (
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
	one [][]int
}

var qs = []question{

	question{
		para{[]int{1, -1, -1, 0}},
		ans{[][]int{
			[]int{-1, 0, 1},
		}},
	},
	question{
		para{[]int{-1, 0, 1, 0}},
		ans{[][]int{
			[]int{-1, 0, 1},
		}},
	},

	question{
		para{[]int{0, 0, 0}},
		ans{[][]int{
			[]int{0, 0, 0},
		}},
	},

	question{
		para{[]int{-2, 0, 1, 1, 2}},
		ans{[][]int{
			[]int{-2, 0, 2},
			[]int{-2, 1, 1},
		}},
	},

	question{
		para{[]int{-2, 0, 0, 2, 2}},
		ans{[][]int{
			[]int{-2, 0, 2},
		}},
	},
	// 如需多个测试，可以复制上方元素。
}

func Test_Problem0015(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
	}
}

func Test_Problem0015_1(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, threeSum1(p.one), "输入:%v", p)
	}
}
