package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lLen := dfs(node.Left) + 1
		rLen := dfs(node.Right) + 1
		ans = max(ans, lLen+rLen)
		return max(lLen, rLen)
	}
	dfs(root)
	return
}
