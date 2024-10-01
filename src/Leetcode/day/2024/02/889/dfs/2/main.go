package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	index := make(map[int]int, n)
	for i, x := range postorder {
		index[x] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(preL, preR, postL, postR int) *TreeNode {
		if preL == preR { // 空节点
			return nil
		}
		if preL+1 == preR { // 叶子节点
			return &TreeNode{preorder[preL], nil, nil}
		}
		leftSize := index[preorder[preL+1]] - postL + 1
		left := dfs(preL+1, preL+1+leftSize, postL, postL+leftSize)
		right := dfs(preL+1+leftSize, preR, postL+leftSize, postR-1)
		return &TreeNode{preorder[preL], left, right}
	}
	return dfs(0, n, 0, n) // 左闭右开区间
}

func main() {

}
