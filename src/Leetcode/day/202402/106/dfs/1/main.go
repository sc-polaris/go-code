package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 { // 空节点
		return nil
	}
	leftSize := slices.Index(inorder, postorder[n-1]) // 左子树的大小
	left := buildTree(inorder[:leftSize], postorder[:leftSize])
	right := buildTree(inorder[leftSize+1:], postorder[leftSize:n-1])
	return &TreeNode{postorder[n-1], left, right}
}

func main() {

}
