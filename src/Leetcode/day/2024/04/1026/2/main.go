package main

import "math"

/*
	归
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt // 保证空间点不影响 mn 和 mx
		}
		lmn, lmx := dfs(node.Left)
		rmn, rmx := dfs(node.Right)
		mn := min(node.Val, lmn, rmn)
		mx := max(node.Val, lmx, rmx)
		ans = max(ans, node.Val-mn, mx-node.Val)
		return mn, mx
	}
	dfs(root)
	return
}
