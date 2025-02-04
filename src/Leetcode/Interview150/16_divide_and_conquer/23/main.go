package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	方法一：分治
*/

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
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
	} else {
		cur.Next = l2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 { // 无需合并，直接返回
		return lists[0]
	}
	left := mergeKLists(lists[:m/2])  // 合并左半部分
	right := mergeKLists(lists[m/2:]) // 合并右半部分
	return mergeTwoLists(left, right) // 左半右半合并
}

/*
	方法二：最小堆
	合并后的第一个节点 first，一定是某个链表的头节点（因为链表已按升序排列）。

	合并后的第二个节点，可能是某个链表的头节点，也可能是 first 的下一个节点。

	例如有三个链表 1->2->5, 3->4->6, 4->5->6，找到第一个节点 1 之后，第二个节点不是另一个链表的头节点，而是节点 1 的下一个节点 2。

	按照这个过程继续思考，每当我们找到一个节点值最小的节点 x，就把节点 x.next 加入「可能是最小节点」的集合中。

	因此，我们需要一个数据结构，它支持：
	· 从数据结构中找到并移除最小节点。
	· 插入节点。
	这可以用最小堆实现。初始把所有链表的头节点入堆，然后不断弹出堆中最小节点 x，如果 x.next 不为空就加入堆中。循环直到堆为空。把弹出的节
	点按顺序拼接起来，就得到了答案。
*/

func mergeKLists2(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)

	dummy := &ListNode{}
	cur := dummy
	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode) // 剩余节点中的最小节点
		if node.Next != nil {            // 下一个节点不为空
			heap.Push(&h, node.Next) // 下一个节点有可能是最小节点，入堆
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next
}

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
