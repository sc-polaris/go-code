package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：前序遍历 递归
func flatten(root *TreeNode) {
	var preorderTraversal func(*TreeNode) []*TreeNode
	preorderTraversal = func(root *TreeNode) []*TreeNode {
		var list []*TreeNode
		if root != nil {
			list = append(list, root)
			list = append(list, preorderTraversal(root.Left)...)
			list = append(list, preorderTraversal(root.Right)...)
		}
		return list
	}
	list := preorderTraversal(root)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

// 方法一：前序遍历 迭代
func flatten2(root *TreeNode) {
	var list []*TreeNode
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			list = append(list, node)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

// 方法二：前序遍历和展开同步进行
func flatten3(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}
		left, right := curr.Left, curr.Right
		if right != nil {
			stack = append(stack, right)
		}
		if left != nil {
			stack = append(stack, left)
		}
		prev = curr
	}
}

// 方法三：寻找前驱节点
func flatten4(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			predecessor.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
