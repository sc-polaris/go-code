package main

/*
	每增加 2 个节点，真二叉树就会多一个叶子，所以一棵有 n 个节点的真二叉树恰好有 (n+1)/2 个叶子
	f[i] 为有 i 个叶子节点的所有真二叉树的列表
	f[1] 为只包含一个节点的二叉树列表
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n&1 == 0 {
		return nil
	}
	f := [11][]*TreeNode{1: {{}}}
	for i := 2; i < len(f); i++ { // 计算 f[i]
		for j := 1; j < i; j++ { // 枚举左子树叶子数
			for _, l := range f[j] { // 枚举左子树
				for _, r := range f[i-j] { // 枚举右子树
					f[i] = append(f[i], &TreeNode{0, l, r})
				}
			}
		}
	}
	return f[(n+1)/2]
}
