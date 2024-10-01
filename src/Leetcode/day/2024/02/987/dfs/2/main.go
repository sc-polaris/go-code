package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	type pair struct{ row, val int }
	groups := map[int][]pair{}
	minCol := 0
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row int, col int) {
		if node == nil {
			return
		}
		groups[col] = append(groups[col], pair{row, node.Val})
		minCol = min(minCol, col)
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	ans := make([][]int, len(groups))
	for i := range ans {
		g := groups[minCol+i]
		slices.SortFunc(g, func(a, b pair) int {
			if a.row != b.row {
				return a.row - b.row
			}
			return a.val - b.val
		})
		ans[i] = make([]int, len(g))
		for j, p := range g {
			ans[i][j] = p.val
		}
	}
	return ans
}

func main() {

}
