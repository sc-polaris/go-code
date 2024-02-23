package main

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	var a []int64
	q := []*TreeNode{root}
	for len(q) > 0 {
		sum := int64(0)
		tmp := q
		q = nil
		for _, node := range tmp {
			sum += int64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		a = append(a, sum)
	}
	n := len(a)
	if k > n {
		return -1
	}
	slices.Sort(a)
	return a[n-k]
}

func main() {

}
