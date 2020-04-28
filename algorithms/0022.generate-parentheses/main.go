package algorithm

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	var ret []string

	helper(n, n, "", &ret)
	return ret
}

func helper(left, right int, s string, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, s)
		return
	}

	if left > 0 {
		helper(left-1, right, s+"(", result)
	}

	if right > 0 && left < right {
		helper(left, right-1, s+")", result)
	}
}
