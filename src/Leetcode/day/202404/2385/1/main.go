package main

/*
	记录父节点 + dfs

	首先从 root 出发这棵树，找到节点值为 start 的节点 startNode。DFS 的同时，用一个哈希表（或者数组）记录每个节点的父节点。

	然后 从 startNode 出发 DFS 这棵树，求出 二叉树的最大深度，即为答案（把 startNode 的深度当成 0）。注意除了递归左右儿子
	以外，还要递归父节点。为避免重复访问节点，可以添加一个递归参数 from，表示当前节点是从 from 过来的，我们不去重复访问节点 form。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	fa := make(map[*TreeNode]*TreeNode)
	var startNode *TreeNode
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(node, from *TreeNode) {
		if node == nil {
			return
		}
		fa[node] = from // 记录每个节点的父节点
		if node.Val == start {
			startNode = node // 找到 start
		}
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	var maxDepth func(*TreeNode, *TreeNode) int
	maxDepth = func(node, from *TreeNode) int {
		if node == nil {
			return -1 // 注意这里是 -1，因为 start 的深度为 0
		}
		res := -1
		if node.Left != from {
			res = max(res, maxDepth(node.Left, node))
		}
		if node.Right != from {
			res = max(res, maxDepth(node.Right, node))
		}
		if fa[node] != from {
			res = max(res, maxDepth(fa[node], node))
		}
		return res + 1
	}
	return maxDepth(startNode, startNode)
}
