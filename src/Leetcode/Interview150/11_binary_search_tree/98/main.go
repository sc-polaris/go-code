package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, left int, right int) bool {
		if node == nil {
			return true
		}
		x := node.Val
		return left < x && x < right && dfs(node.Left, left, x) && dfs(node.Right, x, right)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}

// 中序遍历
func isValidBST2(root *TreeNode) bool {
	pre := math.MinInt
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) || node.Val <= pre {
			return false
		}
		pre = node.Val
		return dfs(node.Right)
	}
	return dfs(root)
}

// 后序遍历
func isValidBST3(root *TreeNode) bool {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt
		}
		lMin, lMax := dfs(node.Left)
		rMin, rMax := dfs(node.Right)
		x := node.Val
		if x <= lMax || x >= rMin {
			return math.MaxInt, math.MaxInt
		}
		return min(lMin, x), max(rMax, x)
	}
	_, mx := dfs(root)
	return mx != math.MaxInt
}
