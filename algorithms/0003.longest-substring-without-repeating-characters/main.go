package algorithm

func lengthOfLongestSubstring(s string) int {
	var count, ret int
	var window [128]int
	for left, right := 0, 0; right < len(s); {
		c := s[right]
		if window[c] == 0 {
			window[c]++
			count++

			right++

			if count > ret {
				ret = count
			}
		} else {
			l := s[left]
			for l != c {
				window[l]--
				count--

				left++
				l = s[left]
			}

			window[l]--
			count--
			left++
		}
	}

	return ret
}
