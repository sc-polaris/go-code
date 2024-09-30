package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 左右子树都找到
	if left != nil && right != nil {
		return root
	}
	// 只有左子树找到
	if left != nil {
		return left
	}
	// 只有右子树找到或者左右子树都没找到
	return right
}
