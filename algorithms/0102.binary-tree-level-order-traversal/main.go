package algorithm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue struct {
	array []*TreeNode
}

func New(size int) *queue {
	return &queue{
		array: make([]*TreeNode, 0, size),
	}
}

func (q *queue) Push(i *TreeNode) {
	q.array = append(q.array, i)
}
func (q *queue) Pop() *TreeNode {
	if q.Len() == 0 {
		return nil
	}

	e := q.array[0]
	q.array = q.array[1:]

	return e
}
func (q *queue) Len() int {
	return len(q.array)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	q := New(10)
	q.Push(root)

	for q.Len() > 0 {
		size := q.Len()
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := q.Pop()
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ret = append(ret, tmp)
	}

	return ret
}
