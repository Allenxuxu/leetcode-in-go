package algorithm

// 滑动窗口
// 左右指针限定窗口
// 移动右边指针扩大窗口，直至找到复合要求字串，记录复合要求的子串
// 然后移动左指针缩小窗口，直至不符合要求，期间如果缩小窗口，字串仍然复合，则记录
// 重复上述步骤，知道窗口滑动到末尾
func minWindow(s string, t string) string {
	var ret string
	need := make(map[uint8]int, len(t))
	have := make(map[uint8]int, len(t))
	for i := range t {
		need[t[i]]++
	}

	var tmp string
	var flag int
	for left, right := 0, 1; right <= len(s); right++ {
		tmp = s[left:right]

		// 包含目标字符串内的字符
		if need[s[right-1]] != 0 {
			have[s[right-1]]++ // 记录下来
			if have[s[right-1]] == need[s[right-1]] {
				flag++
			}
		}

		if len(tmp) < len(t) {
			continue
		}
		// 满足匹配条件了
		if flag == len(need) {
			if len(ret) == 0 || len(tmp) < len(ret) {
				ret = tmp
			}

			// left 指针递增，缩小范围
			for {
				if need[s[left]] != 0 {
					if need[s[left]] == have[s[left]] {
						flag--
					}
					have[s[left]]--
				}

				left++
				tmp = s[left:right]
				if flag != len(need) {
					break
				}
				// 排除子串开头有其他字符 case
				if len(tmp) < len(ret) {
					ret = tmp
				}
			}
		}
	}

	return ret
}
