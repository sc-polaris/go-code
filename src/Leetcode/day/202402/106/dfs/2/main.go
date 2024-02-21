package main

// 时间复杂度O(n) 空间复杂度O(n)

/*
优化：
	1. 用一个哈希表（或者）数组预处理 inorder 每个元素的下标，这样就可以 O(1) 查到
       postorder[n-1] 在 inorder 的位置， 从而 O(1) 知道左子树的大小
	2. 把递归参数改成子数组下标区间（左闭右开区间）的左右端点，从而避免复制数组

如果给定节点的值重复 map[int]int 改成map[int][]int
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	index := make(map[int]int, n)
	for i, x := range inorder {
		index[x] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(inL, inR, postL, postR int) *TreeNode {
		if postL == postR { // 空节点
			return nil
		}
		leftSize := index[postorder[postR-1]] - inL // 左子树的大小
		left := dfs(inL, inL+leftSize, postL, postL+leftSize)
		right := dfs(inL+leftSize+1, inR, postL+leftSize, postR-1)
		return &TreeNode{postorder[postR-1], left, right}
	}
	return dfs(0, n, 0, n) // 左闭右开区间
}

func main() {

}
