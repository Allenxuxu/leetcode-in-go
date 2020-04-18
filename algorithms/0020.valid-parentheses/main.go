package algorithm

import (
	"github.com/Allenxuxu/toolkit/stack"
)

func isValid(s string) bool {
	m := map[string]string{
		"]": "[",
		"}": "{",
		")": "(",
	}

	st := stack.New()
	for _, v := range s {
		element := string(v)
		if _, ok := m[element]; !ok {
			st.Push(element)
		} else {
			if st.Len() == 0 {
				return false
			}
			e := st.Pop()
			if e.(string) != m[element] {
				return false
			}
		}
	}

	return st.Len() == 0
}
