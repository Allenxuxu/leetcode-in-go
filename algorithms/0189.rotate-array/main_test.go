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
	nums []int
	k    int
}

// ans 是答案
type ans struct {
	one []int
}

func Test_rotate1(t *testing.T) {
	ast := assert.New(t)

	var qs = []question{
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6}, 2,
			},
			ans{
				[]int{5, 6, 1, 2, 3, 4},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 14,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 7,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 3,
			},
			ans{
				[]int{5, 6, 7, 1, 2, 3, 4},
			},
		},
	}
	fmt.Printf("rotate1\n")
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		rotate1(p.nums, p.k)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}

func Test_rotate2(t *testing.T) {
	ast := assert.New(t)

	var qs = []question{
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6}, 2,
			},
			ans{
				[]int{5, 6, 1, 2, 3, 4},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 14,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 7,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 3,
			},
			ans{
				[]int{5, 6, 7, 1, 2, 3, 4},
			},
		},
	}
	fmt.Printf("rotate2\n")
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		rotate2(p.nums, p.k)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}

func Test_rotate3(t *testing.T) {
	ast := assert.New(t)

	var qs = []question{
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6}, 2,
			},
			ans{
				[]int{5, 6, 1, 2, 3, 4},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 14,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 7,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},
		question{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 3,
			},
			ans{
				[]int{5, 6, 7, 1, 2, 3, 4},
			},
		},
	}
	fmt.Printf("rotate2\n")
	for _, q := range qs {
		a, p := q.ans, q.para
		fmt.Printf("~~%v~~\n", p)
		rotate3(p.nums, p.k)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}
