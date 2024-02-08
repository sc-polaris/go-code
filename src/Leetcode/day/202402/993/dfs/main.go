package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) (ans bool) {
	depth := 0
	var father *TreeNode
	var dfs func(*TreeNode, *TreeNode, int) bool
	dfs = func(node *TreeNode, fa *TreeNode, d int) bool {
		if node == nil {
			return false
		}
		if node.Val == x || node.Val == y {
			if depth > 0 { // 之前已经找到 x，y 中的一个
				ans = depth == d && father != fa
				return true
			}
			depth, father = d, fa // 之前没找到，记录信息
		}
		return dfs(node.Left, node, d+1) || dfs(node.Right, node, d+1)
	}
	dfs(root, nil, 1)
	return
}

func main() {

}
