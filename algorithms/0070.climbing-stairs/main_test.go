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
type para struct {
	n int
}

// ans 是答案
type ans struct {
	one int
}

var qs = []question{

	question{
		para{
			0,
		},
		ans{
			0,
		},
	},

	question{
		para{
			10,
		},
		ans{
			89,
		},
	},

	question{
		para{
			44,
		},
		ans{
			1134903170,
		},
	},

	question{
		para{
			1,
		},
		ans{
			1,
		},
	},

	// 如需多个测试，可以复制上方元素。
}

func Test_climbStairs(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, climbStairs(p.n), "输入:%v", p)
	}
}

func Test_climbStairs1(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, climbStairs1(p.n), "输入:%v", p)
	}
}

func Test_climbStairs2(t *testing.T) {
	ast := assert.New(t)

	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		ast.Equal(a.one, climbStairs2(p.n), "输入:%v", p)
	}
}
