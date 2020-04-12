package algorithm

// 垂直扫描
func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}

// 水平扫描
func longestCommonPrefix1(strs []string) string {
	short := shortest(strs)

	for i := 0; i < len(strs); i++ {
		for k := range short {
			if strs[i][k] != short[k] {
				short = strs[i][:k]
				break
			}
		}
	}

	return short
}

// 分而治之
// 递归
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	return helper(strs[:len(strs)/2], strs[len(strs)/2:])
}

func helper(l []string, r []string) string {
	if len(l) == 0 {
		return r[0]
	}
	if len(r) == 0 {
		return l[0]
	}

	lStr := helper(l[:len(l)/2], l[len(l)/2:])
	rStr := helper(r[:len(r)/2], r[len(r)/2:])

	return commonPrefix(lStr, rStr)
}

func commonPrefix(l, r string) string {
	length := len(l)
	if len(l) > len(r) {
		length = len(r)
	}
	for i := 0; i < length; i++ {
		if l[i] != r[i] {
			return l[:i]
		}
	}

	return l[:length]
}
