package algorithm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	// tcs is testcase slice
	var tcs = []struct {
		s   []byte
		ans []byte
	}{

		{[]byte("hello"), []byte("olleh")},
		{[]byte("world"), []byte("dlrow")},

		// 可以有多个 testcase
	}

	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		reverseString(tc.s)
		ast.Equal(tc.ans, tc.s, "输入:%v", tc)
	}
}

func Test_reverseString1(t *testing.T) {
	// tcs is testcase slice
	var tcs = []struct {
		s   []byte
		ans []byte
	}{

		{[]byte("hello"), []byte("olleh")},
		{[]byte("world"), []byte("dlrow")},

		// 可以有多个 testcase
	}
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		reverseString1(tc.s)
		ast.Equal(tc.ans, tc.s, "输入:%v", tc)
	}
}
