package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 只在高度相同时匹配

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)
	return max(leftH, rightH) + 1
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // 必须都是 nil
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

func isSubtree(root, subRoot *TreeNode) bool {
	hs := getHeight(subRoot)

	// 返回 node 的高度，以及是否找到了 subRoot
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		leftH, leftOk := dfs(node.Left)
		rightH, rightOk := dfs(node.Right)
		if leftOk || rightOk {
			return 0, true
		}
		nodeH := max(leftH, rightH) + 1
		return nodeH, nodeH == hs && isSameTree(node, subRoot)
	}
	_, ok := dfs(root)
	return ok
}
