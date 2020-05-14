package algorithm

func findAnagrams(s, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	var ret []int
	need := make(map[uint8]int, len(p))
	for i := range p {
		need[p[i]]++
	}

	left, right := 0, len(p)-1

	have := make(map[uint8]int, len(p))

	for i := left; i <= right; i++ {
		if need[s[i]] != 0 {
			have[s[i]]++
		}
	}

	for right < len(s) {
		if match(need, have) {
			ret = append(ret, left)
		}

		left++
		right++
		if right == len(s) {
			break
		}

		have[s[left-1]]--
		if have[s[left-1]] <= 0 {
			delete(have, s[left-1])
		}
		if need[s[right]] != 0 {
			have[s[right]]++
		}
	}

	return ret
}

func match(a, b map[uint8]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range b {
		if a[k] != v {
			return false
		}
	}

	return true
}

func findAnagrams1(s string, p string) []int {
	var res []int
	var need [26]int
	var window [26]int
	var size int
	for i := 0; i < len(p); i++ {
		if need[p[i]-'a'] == 0 {
			size++
		}
		need[p[i]-'a']++
	}
	var left, right int
	var match int
	for right < len(s) {
		// r 是将移入窗口的字符
		r := s[right] - 'a'
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if need[r] > 0 {
			window[r]++
			if window[r] == need[r] {
				match++
			}
		}
		// 判断左侧窗口是否要收缩
		for match == size {
			// 在这里更新最小覆盖子串
			if right-left == len(p) {
				res = append(res, left)
			}
			// l 是将移出窗口的字符
			l := s[left] - 'a'
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			// 左框将向右移，若当前左框字符为符合题意字符，则需要在windowMap中减1
			if need[l] > 0 {
				window[l]--
				// 若当前左边框字符次数 < 该字符目标次数，则匹配的次数match需要 -1
				if window[l] < need[l] {
					match--
				}
			}
		}
	}
	return res
}
