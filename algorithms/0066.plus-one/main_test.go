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

func Test_Problem0066(t *testing.T) {
	ast := assert.New(t)
	var qs = []question{

		question{
			para{[]int{1, 2, 3}},
			ans{[]int{1, 2, 4}},
		},

		question{
			para{[]int{0}},
			ans{[]int{1}},
		},

		question{
			para{[]int{9, 9}},
			ans{[]int{1, 0, 0}},
		},

		question{
			para{[]int{9}},
			ans{[]int{1, 0}},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, plusOne(p.one), "输入:%v", p)
	}
}

func Test_Problem0066_1(t *testing.T) {
	ast := assert.New(t)
	var qs = []question{

		question{
			para{[]int{1, 2, 3}},
			ans{[]int{1, 2, 4}},
		},

		question{
			para{[]int{0}},
			ans{[]int{1}},
		},

		question{
			para{[]int{9, 9}},
			ans{[]int{1, 0, 0}},
		},

		question{
			para{[]int{9}},
			ans{[]int{1, 0}},
		},
	}
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)

		ast.Equal(a.one, plusOne1(p.one), "输入:%v", p)
	}
}
