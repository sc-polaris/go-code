package main

// 时间复杂度O(n) 空间复杂度O(n)

/*
优化：
	1. 用一个哈希表（或者）数组预处理 inorder 每个元素的下标，这样就可以 O(1) 查到
       preorder[0] 在 inorder 的位置， 从而 O(1) 知道左子树的大小
	2. 把递归参数改成子数组下标区间（左闭右开区间）的左右端点，从而避免复制数组

如果给定节点的值重复 map[int]int 改成map[int][]int
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	index := make(map[int]int, n)
	for i, x := range inorder {
		index[x] = i
	}
	var dfs func(int, int, int, int) *TreeNode
	dfs = func(preL, preR, inL, inR int) *TreeNode {
		if preL == preR { // 空节点
			return nil
		}
		leftSize := index[preorder[preL]] - inL
		left := dfs(preL+1, preL+1+leftSize, inL, inL+leftSize)
		right := dfs(preL+1+leftSize, preR, inL+1+leftSize, inR)
		return &TreeNode{preorder[preL], left, right}
	}
	return dfs(0, n, 0, n)
}

func main() {

}
