package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem0126(t *testing.T) {
	ast := assert.New(t)

	// tcs is testcase slice
	tcs := []struct {
		beginWord string
		endWord   string
		wordList  []string
		ans       [][]string
	}{

		{
			"a",
			"c",
			[]string{"a", "b", "c"},
			[][]string{{"a", "c"}},
		},

		{
			"hot",
			"dog",
			[]string{"hot", "dog", "dot"},
			[][]string{
				[]string{"hot", "dot", "dog"},
			},
		},

		{
			"hit",
			"cog",
			[]string{"hot", "dot", "dog", "lot", "log", "cog"},
			[][]string{
				[]string{"hit", "hot", "dot", "dog", "cog"},
				[]string{"hit", "hot", "lot", "log", "cog"},
			},
		},

		// TODO
		//{
		//	"red",
		//	"tax",
		//	[]string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
		//	[][]string{
		//		[]string{"red", "ted", "tad", "tax"},
		//		[]string{"red", "ted", "tex", "tax"},
		//		[]string{"red", "rex", "tex", "tax"},
		//	},
		//},

		// 可以多个 testcase
	}

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)

		ast.Equal(tc.ans, findLadders(tc.beginWord, tc.endWord, tc.wordList), "输入:%v", tc)
	}
}
