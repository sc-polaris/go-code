package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 复制每个节点，把新节点直接插入到原节点的后面
	// 1->1'->2->2'->3->3'
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}

	// 遍历交错链表中的原链表节点
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			// 要复制的 random 是 cur.Random 的下一个节点
			cur.Next.Random = cur.Random.Next
		}
	}

	// 把交错链表分离成两个链表
	newHead := head.Next
	cur := head
	for ; cur.Next.Next != nil; cur = cur.Next {
		clone := cur.Next
		cur.Next = clone.Next        // 恢复原节点的 next
		clone.Next = clone.Next.Next // 设置新节点的 next
	}
	cur.Next = nil
	return newHead
}
