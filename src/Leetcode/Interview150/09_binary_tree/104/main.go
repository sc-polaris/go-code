package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)

	return max(lDepth, rDepth) + 1
}

func maxDepth2(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
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
