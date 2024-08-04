package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 暴力匹配

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // 必须都是 nil
	}

	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) ||
		isSubtree(root.Left, subRoot) ||
		isSubtree(root.Right, subRoot)
}
