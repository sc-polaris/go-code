package main

import (
	"math"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) (ans [][]int) {
	type tuple struct{ col, row, val int }
	var data []tuple
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, row int, col int) {
		if node == nil {
			return
		}
		data = append(data, tuple{col, row, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	slices.SortFunc(data, func(a, b tuple) int {
		if a.col != b.col {
			return a.col - b.col
		}
		if a.row != b.row {
			return a.row - b.row
		}
		return a.val - b.val
	})

	lastCol := math.MinInt
	for _, d := range data {
		if d.col != lastCol {
			lastCol = d.col
			ans = append(ans, []int{})
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], d.val)
	}
	return
}

func main() {

}
