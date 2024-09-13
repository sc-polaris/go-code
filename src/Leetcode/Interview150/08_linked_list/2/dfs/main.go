package dfs

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// l1 和 l2 为当前遍历的节点，carry 为进位
	var addTwo func(l1 *ListNode, l2 *ListNode, carry int) *ListNode
	addTwo = func(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil {
			if carry != 0 {
				return &ListNode{Val: carry} // 如果进位了，就额外创建一个节点
			}
			return nil
		}
		if l1 == nil { // 如果 l1 是空的，那么此时 l2 一定不是空节点
			l1, l2 = l2, l1 // 交换 l1 与 l2，保证 l1 非空，从而简化代码
		}
		sum := carry + l1.Val // 节点值和进位加在一起
		if l2 != nil {
			sum += l2.Val // 节点值和进位加在一起
			l2 = l2.Next  // 下一个节点
		}
		l1.Val = sum % 10                     // 每个节点保存一个数位
		l1.Next = addTwo(l1.Next, l2, sum/10) // 进位
		return l1
	}

	return addTwo(l1, l2, 0)
}
