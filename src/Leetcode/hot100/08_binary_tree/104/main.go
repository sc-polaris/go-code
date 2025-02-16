package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		depth++
		ans = max(ans, depth)
		dfs(node.Left, depth)
		dfs(node.Right, depth)
	}
	dfs(root, 0)
	return
}
