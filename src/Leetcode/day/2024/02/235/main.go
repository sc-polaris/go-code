package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x := root.Val
	if p.Val < x && q.Val < x { // p 和 q 都在左子树
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > x && q.Val > x { // p 和 q 都在右子树
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root // 其它
}

func main() {

}
