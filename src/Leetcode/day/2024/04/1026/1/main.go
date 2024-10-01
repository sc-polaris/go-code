package main

/*
	递
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, mn int, mx int) {
		if node == nil {
			ans = max(ans, mx-mn)
			return
		}
		mn = min(mn, node.Val)
		mx = max(mx, node.Val)
		dfs(node.Left, mn, mx)
		dfs(node.Right, mn, mx)
	}
	dfs(root, root.Val, root.Val)
	return
}
