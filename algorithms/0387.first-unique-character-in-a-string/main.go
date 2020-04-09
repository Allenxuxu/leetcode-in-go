package algorithm

// 时间复杂度 O(n)
// 空间复杂度 O(1)
func firstUniqChar(s string) int {
	// 26 个小写字母
	// 用 map 也可以，但是 数组更快
	rec := make([]int, 26)
	for _, b := range s {
		rec[b-'a']++
	}

	for i, b := range s {
		if rec[b-'a'] == 1 {
			return i
		}
	}

	return -1
}
