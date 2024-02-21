package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	root := &TreeNode{postorder[n-1], nil, nil}
	stk := []*TreeNode{root}
	inorderIndex := len(inorder) - 1
	for i := n - 2; i >= 0; i-- {
		postorderVal := postorder[i]
		node := stk[len(stk)-1]
		if node.Val != inorder[inorderIndex] {
			node.Right = &TreeNode{postorderVal, nil, nil}
			stk = append(stk, node.Right)
		} else {
			for len(stk) > 0 && stk[len(stk)-1].Val == inorder[inorderIndex] {
				node = stk[len(stk)-1]
				stk = stk[:len(stk)-1]
				inorderIndex--
			}
			node.Left = &TreeNode{postorderVal, nil, nil}
			stk = append(stk, node.Left)
		}
	}
	return root
}

func main() {

}
