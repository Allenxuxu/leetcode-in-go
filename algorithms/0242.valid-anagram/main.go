package algorithm

import (
	"sort"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	ss := strings.Split(s, "")
	ts := strings.Split(t, "")

	sort.Strings(ss)
	sort.Strings(ts)

	return strings.EqualFold(strings.Join(ss, ""), strings.Join(ts, ""))
}
