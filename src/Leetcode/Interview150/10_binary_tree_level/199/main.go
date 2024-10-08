package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 先递归右子树，再递归左子树，当某个深度首次到达时，对应的节点就在右视图中。
func rightSideView(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(ans) { // 这个深度首次遇到
			ans = append(ans, node.Val)
		}
		dfs(node.Right, depth+1) // 先递归右子树，保证首次遇到的一定是最右边的节点
		dfs(node.Left, depth+1)
	}
	dfs(root, 0)
	return
}

func rightSideView2(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		for curLen > 0 {
			node := q[0]
			q = q[1:]
			curLen--
			if curLen == 0 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return
}
