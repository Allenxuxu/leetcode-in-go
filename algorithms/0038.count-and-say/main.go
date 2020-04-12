package algorithm

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	last := countAndSay(n - 1)

	return getNextString(last)
}

func getNextString(s string) string {
	if len(s) == 0 {
		return ""
	}

	num := getRepeatNum(s)
	return fmt.Sprintf("%d%c", num, s[0]) + getNextString(s[num:])
}

func getRepeatNum(s string) int {
	var i int
	for i = 1; i < len(s); i++ {
		if s[0] != s[i] {
			break
		}
	}

	return i
}

// 迭代
func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}

	tmp := "1"
	for i := 2; i <= n; i++ {
		tmp = getNextString1(tmp)
	}

	return tmp
}

func getNextString1(s string) string {
	if len(s) == 0 {
		return ""
	}

	var tmp string
	// 统计计数
	var m []struct {
		Value uint8
		Count int
	}
	var index int
	m = append(m, struct {
		Value uint8
		Count int
	}{
		Value: s[0],
		Count: 1,
	})

	for i := 1; i < len(s); i++ {
		if s[i] != m[index].Value {
			m = append(m, struct {
				Value uint8
				Count int
			}{
				Value: s[i],
				Count: 1,
			})
			index++
		} else {
			m[index].Count++
		}
	}

	for _, v := range m {
		tmp += fmt.Sprintf("%d%c", v.Count, v.Value)
	}

	return tmp
}
