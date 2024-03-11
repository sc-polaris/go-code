package main

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Father *TreeNode
}

func inorderSuccessor(p *TreeNode) *TreeNode {
	// Case 1. 如果该节点有右子树，那么下一个节点就是其右子树中最左边的节点
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}

	// Case 3. 如果该节点没有右子树，且是其父节点的右子节点
	// 沿着父指针一直向上，直到找到一个是它父节点的左子节点的节点
	// 如果这样的节点存在，那么这个节点的父节点即是所求
	for p.Father != nil && p == p.Father.Right {
		p = p.Father
	}

	// Case 2. 如果该节点没有右子树，且是其父节点的左子节点，那么下一个节点就是其父节点
	return p.Father
}
