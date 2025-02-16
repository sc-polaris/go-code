package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	leftSize := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
	return &TreeNode{preorder[0], left, right}
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	index := make(map[int]int, n)
	for i, x := range inorder {
		index[x] = i
	}

	var dfs func(preL, preR, inL, inR int) *TreeNode
	dfs = func(preL, preR, inL, inR int) *TreeNode {
		if preL == preR {
			return nil
		}
		leftSize := index[preorder[preL]] - inL
		left := dfs(preL+1, preL+1+leftSize, inL, inL+leftSize)
		right := dfs(preL+1+leftSize, preR, inL+1+leftSize, inR)
		return &TreeNode{preorder[preL], left, right}
	}
	return dfs(0, n, 0, n)
}
