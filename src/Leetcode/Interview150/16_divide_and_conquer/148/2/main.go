package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	归并排序 迭代 改成自底向上计算，空间复杂度优化成 O(1)
	自底向上的意思是：
	· 首先，归并长度为 1 的子链表。例如 [4,2,1,3]，把第一个节点和第二个节点归并，第三个节点和第四个节点归并，得到 [2,4,1,3]。
	· 然后，归并长度为 2 的子链表。例如 [2,4,1,3]，把前两个节点和后两个节点归并，得到 [1,2,3,4]。
	· 然后，归并长度为 4 的子链表。
	· 依此类推，直到归并的长度大于等于链表长度为止，此时链表已经是有序的了。

	具体算法：
	1. 遍历链表，获取链表长度 length。
	2. 初始化步长 step=1。
	3. 循环直到 step≥length。
	4. 每轮循环，从链表头节点开始。
	5. 分割出两段长为 step 的链表，合并，把合并后的链表插到新链表的末尾。重复该步骤，直到链表遍历完毕。
	6. 把 step 扩大一倍。回到第 4 步。
*/

// 获取链表长度
func getListLength(head *ListNode) (res int) {
	for head != nil {
		res++
		head = head.Next
	}
	return
}

// 分割链表
// 如果链表长度 <= size，不做任何操作，返回空节点
// 如果链表长度 > size，把链表的前 size 个节点分割出来（断开连接），并返回剩余链表的头节点
func splitList(head *ListNode, size int) *ListNode {
	// 先找到 nextHead 的前一个节点
	cur := head
	for i := 0; i < size-1 && cur != nil; i++ {
		cur = cur.Next
	}

	// 如果链表长度 <= size
	if cur == nil || cur.Next == nil {
		return nil // 不做任何操作，返回空节点
	}

	nextHead := cur.Next
	cur.Next = nil // 断开 nextHead 的前一个节点和 nextHead 的连接
	return nextHead
}

// 返回合并后的链表的头节点和尾节点
func mergeTwoLists(l1, l2 *ListNode) (head, tail *ListNode) {
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

	for cur.Next != nil {
		cur = cur.Next
	}

	return dummy.Next, cur
}

func sortList(head *ListNode) *ListNode {
	length := getListLength(head)
	dummy := ListNode{Next: head}
	// step 为步长，即参与合并的链表长度
	for step := 1; step < length; step *= 2 {
		newListTail := &dummy // 新链表的末尾
		cur := dummy.Next     // 每轮循环的起始节点
		for cur != nil {
			// 从 cur 开始，分割出两段长为 step 的链表，头节点分别为 head1 和 head2
			head1 := cur
			head2 := splitList(head1, step)
			cur = splitList(head2, step) // 下一轮循环的起始节点
			// 合并两段长尾 step 的链表
			head, tail := mergeTwoLists(head1, head2)
			// 合并后的头节点 head，插到 newListTail 的后面
			newListTail.Next = head
			newListTail = tail // tail 现在是新链表的末尾
		}
	}
	return dummy.Next
}
