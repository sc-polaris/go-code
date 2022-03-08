package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func findKthToTail(pListHead *ListNode, k int) *ListNode {
	var n int
	for p := pListHead; p != nil; p = p.Next {
		n++
	}
	if n < k {
		return nil
	}
	var p = pListHead
	for i := 0; i < n-k; i++ {
		p = p.Next
	}
	return p
}
