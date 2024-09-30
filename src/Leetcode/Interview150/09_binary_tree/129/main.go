package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs 无返回值
func sumNumbers(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, x int) {
		if node == nil {
			return
		}
		x = x*10 + node.Val
		if node.Left == node.Right { // 叶子节点
			ans += x
			return
		}
		dfs(node.Left, x)
		dfs(node.Right, x)
	}
	dfs(root, 0)
	return
}

// dfs 2
func sumNumbers2(root *TreeNode) (ans int) {
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, x int) int {
		if node == nil {
			return 0
		}
		x = x*10 + node.Val
		if node.Left == node.Right {
			return x
		}
		return dfs(node.Left, x) + dfs(node.Right, x)
	}
	return dfs(root, 0)
}
