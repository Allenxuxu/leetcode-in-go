package algorithm

import (
	"fmt"

	"github.com/Allenxuxu/toolkit/queue"
)

// TODO 未完成
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var ret [][]string

	m := make(map[string]bool, len(wordList))
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = true
	}
	m[beginWord] = false

	q := queue.New()
	q.Push(beginWord)

	path := make(map[string][]string)
	path[beginWord] = make([]string, 0)
	for q.Len() > 0 {
		size := q.Len()
		for k := 0; k < size; k++ {
			word := q.Pop().(string)

			for i := 0; i < len(word); i++ {
				bWord := []byte(word)
				for j := 0; j < 26; j++ {
					b := byte('a' + j)
					if bWord[i] == b {
						continue
					}

					bWord[i] = b
					if m[string(bWord)] {
						s := path[word]
						s = append(s, string(bWord))
						path[word] = s

						if string(bWord) == endWord {
							break
						}

						q.Push(string(bWord))
						m[string(bWord)] = false
					}
				}
			}
		}
	}
	var s []string
	s = append(s, beginWord)
	fmt.Println(path)
	dfs(beginWord, path, &s, endWord, &ret)

	for i := 0; i+1 < len(ret); i++ {
		if len(ret[i]) > len(ret[i+1]) {
			ret = append(ret[:i], ret[i+1:]...)
		} else if len(ret[i]) < len(ret[i+1]) {
			ret = append(ret[:i+1], ret[i+2:]...)
		}
	}
	return ret
}

func dfs(current string, path map[string][]string, s *[]string, endWord string, ret *[][]string) {
	if current == endWord {
		tmp := make([]string, len(*s))
		copy(tmp, *s)
		*ret = append(*ret, tmp)
		return
	}

	for _, v := range path[current] {
		*s = append(*s, v)
		dfs(v, path, s, endWord, ret)
		*s = (*s)[:len(*s)-1]
	}
}
