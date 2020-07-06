package algorithm

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 || (str[0] != '+' && str[0] != '-' && !unicode.IsDigit(rune(str[0]))) {
		return 0
	}

	var pos = 1
	if str[0] == '-' {
		pos = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	index := strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	if index == -1 {
		index = len(str)
	}

	s := str[:index]

	var ret int
	for i := 0; i < len(s); i++ {
		num := int(s[i] - '0')
		// 2147483647， -2147483648
		// 如果ret是正数， ret + num > 2147483647(math.MaxInt32), 那么如果是负数其绝对值 >= 2147483648 , 所以无需额外判断 math.MinInt32
		if ret > math.MaxInt32/10 || (ret == math.MaxInt32/10 && num > 7) {
			if pos == -1 {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		ret = ret*10 + num
	}

	return ret * pos
}
