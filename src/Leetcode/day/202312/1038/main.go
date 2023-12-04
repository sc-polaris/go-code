package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	val := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right) // 递归右子树
		val += node.Val
		node.Val = val
		dfs(node.Left) // 递归左子树
	}
	dfs(root)
	return root
}

func bstToGst2(root *TreeNode) *TreeNode {
	val := 0
	node := root
	for node != nil {
		if node.Right == nil {
			val += node.Val
			node.Val = val
			node = node.Left
		} else {
			succ := getSuccessor(node)
			if succ.Left == nil {
				succ.Left = node
				node = node.Right
			} else {
				succ.Left = nil
				val += node.Val
				node.Val = val
				node = node.Left
			}
		}
	}
	return root
}

func getSuccessor(node *TreeNode) *TreeNode {
	succ := node.Right
	for succ.Left != nil && succ.Left != node {
		succ = succ.Left
	}
	return succ
}

func main() {

}
