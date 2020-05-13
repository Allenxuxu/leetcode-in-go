package algorithm

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

		if need[s[right-1]] != 0 {
			have[s[right-1]]++
			if have[s[right-1]] == need[s[right-1]] {
				flag++
			}
		}

		if len(tmp) < len(t) {
			continue
		}
		if flag == len(need) {
			if len(ret) == 0 || len(tmp) < len(ret) {
				ret = tmp
			}

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

				if len(tmp) < len(ret) {
					ret = tmp
				}
			}
		}
	}

	return ret
}
