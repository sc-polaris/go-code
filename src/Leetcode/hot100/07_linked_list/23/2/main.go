package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 分治

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
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
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	l := mergeKLists(lists[:m/2]) // 合并左半部分
	r := mergeKLists(lists[m/2:]) // 合并有半部分
	return mergeTwoLists(l, r)
}
