package main

/*
	 分类讨论：
		当前节点是空节点，当前节点是 p，当前节点是 q，返回当前节点
		左右子树都找到，返回当前节点
		只有左子树找到，返回递归左子树的结果
		只有右子树找到，返回递归右子树的结果
		左右子树都没有找到，返回空节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func main() {

}
