package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	假设进环前的路程为 a，环长为 b。设慢指针走了 x 步时，快慢指针相遇，此时快指针走了 2x步。
	2x-x = nb （快指针比慢指针多走了 n 圈）
	   x = nb	慢指针总共走过的路程是 nb

	但这 nb 当中，实际上包含了进环前的一个小 a，因此慢指针在环中只走了 nb-a 步，它还得再往前走 a 步，才是完整的 n 圈。
	所以，我们让头节点和慢指针同时往前走，当他俩相遇时，就走过了最后这 a 步。
*/

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}
