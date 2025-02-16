package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) (ans int) {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		k--
		if k == 0 {
			ans = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return
}
