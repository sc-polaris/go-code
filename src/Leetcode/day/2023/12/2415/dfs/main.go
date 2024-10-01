package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode, *TreeNode, bool)
	dfs = func(x *TreeNode, y *TreeNode, isOdd bool) {
		if x == nil {
			return
		}
		if isOdd {
			x.Val, y.Val = y.Val, x.Val
		}
		dfs(x.Left, y.Right, !isOdd)
		dfs(x.Right, y.Left, !isOdd)
	}

	dfs(root.Left, root.Right, true)

	return root
}

func main() {

}
