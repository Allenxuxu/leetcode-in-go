package algorithm

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func reverseString1(s []byte) {
	i := 0
	j := len(s) - 1

	swap(s, i, j)
}

func swap(s []byte, i, j int) {
	if i < j {
		s[i], s[j] = s[j], s[i]
		swap(s, i+1, j-1)
	}
}
