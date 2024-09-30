package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)   // 翻转左子树
	right := invertTree(root.Right) // 翻转右子树
	root.Left = right
	root.Right = left
	return root
}
