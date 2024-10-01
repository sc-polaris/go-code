package main

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

	root := &TreeNode{preorder[0], nil, nil}
	stk := []*TreeNode{root}
	inorderIndex := 0
	for i := 1; i < n; i++ {
		preorderVal := preorder[i]
		node := stk[len(stk)-1]
		if node.Val != inorder[inorderIndex] {
			node.Left = &TreeNode{preorderVal, nil, nil}
			stk = append(stk, node.Left)
		} else {
			for len(stk) > 0 && stk[len(stk)-1].Val == inorder[inorderIndex] {
				node = stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				inorderIndex++
			}
			node.Right = &TreeNode{preorderVal, nil, nil}
			stk = append(stk, node.Right)
		}
	}

	return root
}

func main() {

}
