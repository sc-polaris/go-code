package main

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func getTail(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func quickSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, mid, right := &ListNode{}, &ListNode{}, &ListNode{}
	leftTail, midTail, rightTail := left, mid, right
	val := head.Val

	for p := head; p != nil; p = p.Next {
		if p.Val < val {
			leftTail.Next = p
			leftTail = leftTail.Next
		} else if p.Val == val {
			midTail.Next = p
			midTail = midTail.Next
		} else {
			rightTail.Next = p
			rightTail = rightTail.Next
		}
	}

	leftTail.Next = nil
	midTail.Next = nil
	rightTail.Next = nil

	left.Next = quickSortList(left.Next)
	right.Next = quickSortList(right.Next)

	getTail(left).Next = mid.Next
	getTail(left).Next = right.Next

	return left.Next
}
