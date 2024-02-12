package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (ans []int) {
	var stk []*TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			ans = append(ans, root.Val)
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1].Right
		stk = stk[:len(stk)-1]
	}
	return
}

func main() {

}
