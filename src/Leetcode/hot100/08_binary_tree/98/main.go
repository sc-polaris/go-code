package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left int, right int) bool
	dfs = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}
		x := node.Val
		return left < x && x < right &&
			dfs(node.Left, left, x) &&
			dfs(node.Right, x, right)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
