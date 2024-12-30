package main

/*
	给你一棵以 root 为根的二叉树和一个 head 为第一个节点的链表。

	如果在二叉树中，存在一条一直向下的路径，且每个点的数值恰好一一对应以 head 为首的链表中每个节点的值，那么请你返回 True ，否则返回 False 。

	一直向下的路径的意思是：从树中某个节点开始，一直连续向下的路径。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(s *ListNode, t *TreeNode) bool {
		if s == nil { // 整个链表匹配完毕
			return true
		}
		// 否则需要继续匹配
		if t == nil { // 无法继续匹配
			return false
		}
		return s.Val == t.Val && (dfs(s.Next, t.Left) || dfs(s.Next, t.Right)) ||
			s == head && (dfs(head, t.Left) || dfs(head, t.Right))
	}
	return dfs(head, root)
}
