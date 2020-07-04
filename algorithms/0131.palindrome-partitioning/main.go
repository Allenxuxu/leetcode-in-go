package algorithm

func partition(s string) [][]string {
	var ret [][]string
	backtracking(s, 0, &[]string{}, &ret)
	return ret
}

func backtracking(s string, start int, res *[]string, ret *[][]string) {
	if start == len(s) {
		tmp := make([]string, len(*res))
		copy(tmp, *res)
		*ret = append(*ret, tmp)

		return
	}

	for i := start; i < len(s); i++ {
		if !isPalindrome(s, start, i) {
			continue
		}

		*res = append(*res, s[start:i+1])
		backtracking(s, i+1, res, ret)
		*res = (*res)[:len(*res)-1]
	}
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
