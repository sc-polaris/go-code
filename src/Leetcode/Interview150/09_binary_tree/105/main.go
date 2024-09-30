package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	index := make(map[int]int)
	for i, x := range inorder {
		index[x] = i
	}
	var dfs func(int, int, int, int) *TreeNode
	dfs = func(preL, preR, inL, inR int) *TreeNode {
		if preL == preR {
			return nil
		}
		leftSize := index[preorder[preL]] - inL
		left := dfs(preL+1, preL+1+leftSize, inL, inL+leftSize)
		right := dfs(preL+1+leftSize, preR, inL+1+leftSize, inR)
		return &TreeNode{preorder[preL], left, right}
	}
	return dfs(0, n, 0, n) // 左闭右开区间
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	stack := []*TreeNode{root}
	inorderIndex := 0
	for i := 1; i < n; i++ {
		preorderVal := preorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{Val: preorderVal}
			stack = append(stack, node.Left)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{Val: preorderVal}
			stack = append(stack, node.Right)
		}
	}
	return root
}
