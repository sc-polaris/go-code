package main

import (
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 { // 空节点
		return nil
	}
	if n == 1 { // 叶子节点
		return &TreeNode{preorder[0], nil, nil}
	}
	leftSize := slices.Index(postorder, preorder[1]) + 1 // 左子树的大小
	left := constructFromPrePost(preorder[1:1+leftSize], postorder[:leftSize])
	right := constructFromPrePost(preorder[1+leftSize:], postorder[leftSize:n-1])
	return &TreeNode{preorder[0], left, right}
}

func main() {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	constructFromPrePost(preorder, postorder)
}
