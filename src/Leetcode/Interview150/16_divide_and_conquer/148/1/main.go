package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序 分治

// middleNode 链表的中间结点（快慢指针）
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next // 下一个节点就是链表的中间结点 mid
	slow.Next = nil  // 断开 mid 的前一个节点和 mid 的连接
	return mid
}

// 合并两个有序链表（双指针）
func mergeTowList(l1, l2 *ListNode) *ListNode {
	dummy := ListNode{} // 哨兵节点
	cur := &dummy       // cur 指向新链表的结尾
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	// 拼接剩余的链表
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head2 := middleNode(head)
	head = sortList(head)
	head2 = sortList(head2)
	return mergeTowList(head, head2)
}
