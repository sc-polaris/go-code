package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	countLevel := func(node *TreeNode) (ans int) {
		for node != nil {
			ans++
			node = node.Left
		}
		return
	}

	left := countLevel(root.Left)
	right := countLevel(root.Right)
	if left == right {
		return countNodes(root.Right) + (1 << left)
	}
	return countNodes(root.Left) + (1 << right)
}
