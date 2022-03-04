package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func entryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	first, second := head, head
	for first != nil && second != nil {
		first = first.Next
		second = second.Next
		if second != nil {
			second = second.Next
		} else {
			return nil
		}

		if first == second {
			first = head
			for first != second {
				first = first.Next
				second = second.Next
			}
			return first
		}
	}

	return nil
}
