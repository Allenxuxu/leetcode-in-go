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
