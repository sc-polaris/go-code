package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}

	dummy := &ListNode{Next: head}
	p0 := dummy

	var pre, cur *ListNode = nil, p0.Next
	for ; n >= k; n -= k {
		for range k {
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}
