package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// dfs
func connect(root *Node) *Node {
	var pre []*Node
	var dfs func(*Node, int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}
		if depth == len(pre) {
			pre = append(pre, node)
		} else {
			pre[depth].Next = node
			pre[depth] = node
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return root
}

// bfs
func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}
	for q != nil {
		tmp := q
		q = nil
		for i, node := range tmp {
			if i > 0 { // 连接同一层的两个相邻节点
				tmp[i-1].Next = node
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

// bfs + 链表
/*
	从第一层开始（第一层只有一个 root 节点），每次循环：
	遍历当前层的链表节点，通过节点的 left 和 right 得到下一层的节点。
	把下一层的节点从左到右连接成一个链表。
	拿到下一层链表的头节点，进入下一轮循环。
*/

func connect3(root *Node) *Node {
	dummy := &Node{}
	cur := root
	for cur != nil {
		dummy.Next = nil
		next := dummy    // 下一层的链表
		for cur != nil { // 遍历当前层的链表
			if cur.Left != nil {
				next.Next = cur.Left // 下一层的相邻节点连起来
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
