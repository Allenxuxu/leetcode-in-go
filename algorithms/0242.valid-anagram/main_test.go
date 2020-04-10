package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans bool
}{

	{"anagram", "nagaram", true},
	{"rat", "car", false},
	{"at", "car", false},
	{"", "", true},
	{"世界和平", "和平世界", true},
	{"ac", "bb", false},
}

func Test_isAnagram(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isAnagram(tc.s, tc.t), "输入:%v", tc)
	}
}

func Test_isAnagram1(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isAnagram1(tc.s, tc.t), "输入:%v", tc)
	}
}
