package algorithm

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	var m = make(map[string][]string)
	for _, v := range strs {
		tmp := strings.Split(v, "")
		sort.Strings(tmp)
		key := strings.Join(tmp, "")
		m[key] = append(m[key], v)
	}

	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func groupAnagrams1(strs []string) [][]string {
	var ret [][]string
	// golang 可以直接比较 数组的 （切片不可以）
	var m = make(map[[26]byte][]string)
	for _, v := range strs {
		key := hashCode(v)
		m[key] = append(m[key], v)
	}

	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func hashCode(s string) [26]byte {
	var code [26]byte
	for i := 0; i < len(s); i++ {
		code[s[i]-'a']++
	}
	return code
}

// ** golang 可以直接比较 数组的 （切片不可以）
//func main() {
//	var a [2]int = [2]int{1, 2}
//	var b [2]int
//	print(a == b)
//}
