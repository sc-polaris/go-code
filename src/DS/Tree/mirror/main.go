package mirror

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	Morris 实现前序遍历 根左右
	 1. 若当前节点 root 的左子树为空，将当前节点值添加至结果列表 ans 中，并将当前节点更新为 root.right
	 2. 若当前节点 root 的左子树不为空，找到左子树的最右节点 pre（也即是 root 节点在中序遍历下的前驱节点）：
		2.1 若前驱节点 pre 的右子树为空，将当前节点值添加至结果列表 ans 中，然后将前驱节点的右子树指向当前节点 root，并将当前节点更新为 root.left。
		2.2 若前驱节点 pre 的右子树不为空，将前驱节点右子树指向空（即解除 pre 与 root 的指向关系），并将当前节点更新为 root.right。

	循环以上步骤，直至二叉树节点为空，遍历结束。
*/

func preorderTraversal(root *TreeNode) (ans []int) {
	for root != nil {
		if root.Left == nil {
			ans = append(ans, root.Val)
			root = root.Right
		} else {
			prev := root.Left
			for prev.Right != nil && prev.Right != root {
				prev = prev.Right
			}
			if prev.Right == nil {
				ans = append(ans, root.Val)
				prev.Right = root
				root = root.Left
			} else {
				prev.Right = nil
				root = root.Right
			}
		}
	}
	return
}

/*
	Morris 实现中序遍历 左根右
	1. 若当前节点 root 的左子树为空，将当前节点值添加至结果列表 ans 中，并将当前节点更新为 root.right
	2. 若当前节点 root 的左子树不为空，找到左子树的最右节点 prev（也即是 root 节点在中序遍历下的前驱节点）：
		2.1 若前驱节点 prev 的右子树为空，将前驱节点的右子树指向当前节点 root，并将当前节点更新为 root.left。
		2.2 若前驱节点 prev 的右子树不为空，将当前节点值添加至结果列表 ans 中，然后将前驱节点右子树指向空（即解除 prev 与 root 的指向关系），并将当前节点更新为 root.right。
	循环以上步骤，直至二叉树节点为空，遍历结束。
*/

func inorderTraversal(root *TreeNode) (ans []int) {
	for root != nil {
		if root.Left == nil {
			ans = append(ans, root.Val)
			root = root.Right
		} else {
			prev := root.Left
			for prev.Right != nil && prev.Right != root {
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

/*
	Morris 实现后序遍历
	 1. 若当前节点 root 的右子树为空，将当前节点值添加至结果列表 ans 中，并将当前节点更新为 root.left
	 2. 若当前节点 root 的右子树不为空，找到右子树的最左节点 next（也即是 root 节点在中序遍历下的后继节点）：
		2.1 若后继节点 next 的左子树为空，将当前节点值添加至结果列表 ans 中，然后将后继节点的左子树指向当前节点 root，并将当前节点更新为 root.right。
		2.2 若后继节点 next 的左子树不为空，将后继节点左子树指向空（即解除 next 与 root 的指向关系），并将当前节点更新为 root.left。

	循环以上步骤，直至二叉树节点为空，遍历结束。
	最后返回结果列表的逆序即可。

	Morris 后序遍历跟 Morris 前序遍历思路一致，只是将前序的“根左右”变为“根右左”，最后逆序结果即可变成“左右根”。
*/

func postorderTraversal(root *TreeNode) (ans []int) {
	for root != nil {
		if root.Right == nil {
			ans = append([]int{root.Val}, ans...)
			root = root.Left
		} else {
			next := root.Right
			for next.Left != nil && next.Left != root {
				next = next.Left
			}
			if next.Left == nil {
				ans = append([]int{root.Val}, ans...)
				next.Left = root
				root = root.Right
			} else {
				next.Left = nil
				root = root.Left
			}
		}
	}
	return
}
