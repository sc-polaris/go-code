package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// r比l多n个节点

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	l, r := dummy, dummy
	for ; n > 0; n-- {
		r = r.Next
	}
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
