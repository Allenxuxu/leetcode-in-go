package algorithm

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	need := make(map[uint8]int, len(s1))
	for i := range s1 {
		need[s1[i]]++
	}

	left, right := 0, len(s1)-1

	have := make(map[uint8]int, len(s1))

	for i := left; i <= right; i++ {
		if need[s2[i]] != 0 {
			have[s2[i]]++
		}
	}

	for right < len(s2) {
		if match(need, have) {
			return true
		}

		left++
		right++
		if right == len(s2) {
			break
		}

		have[s2[left-1]]--
		if have[s2[left-1]] <= 0 {
			delete(have, s2[left-1])
		}
		if need[s2[right]] != 0 {
			have[s2[right]]++
		}
	}

	return false
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
