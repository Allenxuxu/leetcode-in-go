package algorithm

type queue struct {
	array []string
}

func New(size int) *queue {
	return &queue{
		array: make([]string, 0, size),
	}
}

func (q *queue) Push(i string) {
	q.array = append(q.array, i)
}

func (q *queue) Pop() string {
	if q.Len() == 0 {
		return ""
	}

	e := q.array[0]
	q.array = q.array[1:]

	return e
}

func (q *queue) Len() int {
	return len(q.array)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := make(map[string]bool, len(wordList))
	for i := 0; i < len(wordList); i++ {
		m[wordList[i]] = true
	}

	m[beginWord] = false

	q := New(len(wordList))
	q.Push(beginWord)
	count := 1
	for q.Len() > 0 {
		size := q.Len()

		for i := 0; i < size; i++ {
			word := []byte(q.Pop())
			for i := 0; i < len(word); i++ {
				tmp := word[i]

				for j := 0; j < 26; j++ {
					b := byte('a' + j)
					if b == word[i] {
						continue
					}
					word[i] = b
					if m[string(word)] {
						if string(word) == endWord {
							return count + 1
						}

						q.Push(string(word))

						m[string(word)] = false
					}
				}
				word[i] = tmp
			}
		}
		count++
	}

	return 0
}
