package algorithm

import "strings"

func isPalindrome(s string) bool {
	// O(n)
	s = strings.ToLower(s)

	var (
		i int
		j = len(s) - 1
	)

	for i < j {
		if !check(s[i]) {
			i++
			continue
		}
		if !check(s[j]) {
			j--
			continue
		}

		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func check(e byte) bool {
	if (e < 'a' || e > 'z') && (e < '0' || e > '9') {
		return false
	}

	return true
}
