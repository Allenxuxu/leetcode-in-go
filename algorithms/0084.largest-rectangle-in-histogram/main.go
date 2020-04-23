package algorithm

type stack struct {
	array []int
}

func New(size int) *stack {
	return &stack{
		array: make([]int, 0, size),
	}
}

func (s *stack) Push(i int) {
	s.array = append(s.array, i)
}

func (s *stack) Pop() int {
	if len(s.array) == 0 {
		return 0
	}

	e := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return e
}

func (s *stack) Peek() int {
	if len(s.array) == 0 {
		return 0
	}
	return s.array[len(s.array)-1]
}

func (s *stack) Len() int {
	return len(s.array)
}

// 单调栈解法
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var ans int
	s := New(len(heights) / 2)
	heights = append(heights[:1], heights...)
	heights[0] = 0
	heights = append(heights, 0)
	for i := 0; i < len(heights); i++ {
		for s.Len() != 0 && heights[s.Peek()] > heights[i] {
			cur := s.Pop()
			left := s.Peek() + 1
			right := i - 1

			ans = max(ans, heights[cur]*(right-left+1))
		}

		s.Push(i)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
