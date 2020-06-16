package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

// 无法找到前面的节点，只能将后面一个节点复制到node，然后删除后一个节点
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
