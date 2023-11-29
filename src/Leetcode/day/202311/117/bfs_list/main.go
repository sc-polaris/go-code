package bfs_list

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	dummy := &Node{} // 哨兵节点，第一个节点之前的节点
	cur := root
	for cur != nil {
		dummy.Next = nil
		next := dummy    // 下一层的链表
		for cur != nil { // 遍历当前层的链表
			if cur.Left != nil {
				next.Next = cur.Left // 下一程的相邻节点连起来
				next = cur.Left
			}
			if cur.Right != nil {
				next.Next = cur.Right // 下一层的相邻节点连起来
				next = cur.Right
			}
			cur = cur.Next // 当前层链表的下一个节点
		}
		cur = dummy.Next // 下一层链表的头节点
	}
	return root
}
