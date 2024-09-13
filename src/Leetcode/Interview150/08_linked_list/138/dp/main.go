package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
	我们首先将该链表中每一个节点拆分为两个相连的节点，例如对于链表 A→B→C，我们可以将其拆分为 A→A′→B→B′→C→C′。
	对于任意一个原节点 S，其拷贝节点 S′ 即为其后继节点。

	这样，我们可以直接找到每一个拷贝节点 S' 的随机指针应当指向的节点，即为其原节点 S 的随机指针指向的节点 T 的后继节点
	T'。需要注意原节点的随机指针可能为空，我们需要特别判断这种情况。

	当我们完成了拷贝节点的随机指针的赋值，我们只需要将这个链表按照原节点与拷贝节点的种类进行拆分即可，只需要遍历一次。同样
	需要注意最后一个拷贝节点的后继节点为空，我们需要特别判断这种情况。

*/

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next
		}
	}

	return headNew
}
