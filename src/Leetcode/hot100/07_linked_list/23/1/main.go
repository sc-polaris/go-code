package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 最小堆

func mergeKLists(lists []*ListNode) *ListNode {
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
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
