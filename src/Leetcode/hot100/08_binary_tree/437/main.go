package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) (ans int) {
	cnt := map[int]int{0: 1}
	var dfs func(node *TreeNode, s int)
	dfs = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		s += node.Val
		ans += cnt[s-targetSum]
		cnt[s]++
		dfs(node.Left, s)
		dfs(node.Right, s)
		cnt[s]-- // 恢复现场
	}
	dfs(root, 0)
	return
}
