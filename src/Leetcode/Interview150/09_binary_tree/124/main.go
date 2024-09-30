package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0 // 没有节点和为 0
		}
		lVal := dfs(node.Left)                  // 左子树的最大链和
		rVal := dfs(node.Right)                 // 右子树的最大链和
		ans = max(ans, lVal+rVal+node.Val)      // 两条链拼成路径
		return max(max(lVal, rVal)+node.Val, 0) // 当前子树最大链和
	}
	dfs(root)
	return ans
}
