package algorithm

import (
	"strings"
)

func minWindow(s string, t string) string {
	var ret string
	need := make(map[uint8]int, len(t))
	for i := range t {
		need[t[i]]++
	}

	var tmp string
	for left, right := 0, 1; right <= len(s); right++ {
		tmp = s[left:right]

		if len(tmp) < len(t) {
			continue
		}

		if valid(tmp, need) {
			if len(ret) == 0 || len(tmp) < len(ret) {
				ret = tmp
			}

			for {
				left++
				tmp = s[left:right]
				if !valid(tmp, need) {
					break
				}
				if len(tmp) < len(ret) {
					ret = tmp
				}
			}
		}
	}

	return ret
}

func valid(s string, m map[uint8]int) bool {
	for k, v := range m {
		if strings.Count(s, string(k)) < v {
			return false
		}
	}

	return true
}
