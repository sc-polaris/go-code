package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	cur := head
	for ; cur.Next != nil; cur = cur.Next {
		n++
	}

	cur.Next = head // 闭环

	//新链表的尾部
	for i := 0; i < (n - k%n); i++ {
		cur = cur.Next
	}

	// 新链表的头部
	headNew := cur.Next
	cur.Next = nil // 断开
	return headNew
}
