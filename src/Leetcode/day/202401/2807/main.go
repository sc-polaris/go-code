package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for cur := head; cur.Next != nil; cur = cur.Next.Next {
		cur.Next = &ListNode{gcd(cur.Val, cur.Next.Val), cur.Next}
	}
	return head
}

func main() {

}
