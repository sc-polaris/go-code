package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (ans []int) {
	for root != nil {
		if root.Left == nil { // 没有左子树
			ans = append(ans, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		} else { // 有左子树
			// 找到前驱节点
			prev := root.Left
			for prev.Right != nil && prev.Right != root {
				// 有右子树且没有设置过指向 root，则继续
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = root
				root = root.Left
			} else {
				ans = append(ans, root.Val)
				prev.Right = nil
				root = root.Right
			}
		}
	}
	return
}

func main() {

}
