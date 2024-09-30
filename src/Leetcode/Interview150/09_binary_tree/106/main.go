package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	index := make(map[int]int)
	for i, x := range inorder {
		index[x] = i
	}

	var dfs func(int, int, int, int) *TreeNode
	dfs = func(inL, inR, postL, postR int) *TreeNode {
		if postL == postR {
			return nil
		}
		leftSize := index[postorder[postR-1]] - inL // 左子树的大小
		left := dfs(inL, inL+leftSize, postL, postL+leftSize)
		right := dfs(inL+leftSize+1, inR, postL+leftSize, postR-1)
		return &TreeNode{postorder[postR-1], left, right}
	}

	return dfs(0, n, 0, n) // 左闭右开区间
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[n-1]}
	stack := []*TreeNode{root}
	inorderIndex := n - 1
	for i := n - 2; i >= 0; i-- {
		postorderVal := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{Val: postorderVal}
			stack = append(stack, node.Left)
		}
	}
	return root
}
