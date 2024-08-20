package Iter

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历的顺序是：根节点 -> 左子树 -> 右子树。迭代实现时，我们可以先访问根节点，然后将右子节点（如果有的话）先于左子节点入栈，这样出栈时就能保证左子树先被访问。
func preorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}

// 中序遍历的顺序是：左子树 -> 根节点 -> 右子树。迭代现的关键在于先将所有的左子节点压入栈中，直到遇到空节点开始弹出并访问节点，然后转向其右子节点。
func inorderTraversal(root *TreeNode) (ans []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, root.Val)
			root = root.Right
		}
	}
	return
}

// 后序遍历的顺序是：左子树 -> 右子树 -> 根节点。 其中关键点在于要记录上一次输出的节点
func postorderTraversal(root *TreeNode) (ans []int) {
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Right == nil || root.Right == prev {
				ans = append(ans, root.Val)
				prev, root = root, nil
			} else {
				stack = append(stack, root)
				root = root.Right
			}
		}
	}
	return
}

// 相当于类似的前序翻转
//func postorderTraversal(root *TreeNode) (ans []int) {
//	if root == nil {
//		return nil
//	}
//	stack := []*TreeNode{root}
//	for len(stack) > 0 {
//		node := stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		ans = append(ans, node.Val)
//		if node.Left != nil {
//			stack = append(stack, node.Left)
//		}
//		if node.Right != nil {
//			stack = append(stack, node.Right)
//		}
//	}
//	slices.Reverse(ans)
//	return
//}
