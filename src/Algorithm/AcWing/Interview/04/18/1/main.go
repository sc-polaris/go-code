package _

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
具体步骤如下：
	1. 先利用前序遍历找根节点：前序遍历的第一个数，就是根节点的值；
	2. 在中序遍历中找到根节点的位置k，则k左边是左子树的中序遍历，右边是右子树的中序遍历；
	3. 假设左子树的中序遍历的长度是l，则在前序遍历中，根节点后面的l个数，是左子树的前序遍历，剩下的数是右子树的前序遍历；
	4. 有了左右子树的前序遍历和中序遍历，我们可以先递归创建出左右子树，然后再创建根节点；
*/

var pos = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	for i := 0; i < n; i++ {
		pos[inorder[i]] = i
	}
	return dfs(preorder, inorder, 0, n-1, 0, n-1)
}

func dfs(pre, in []int, pl, pr, il, ir int) *TreeNode {
	if pl > pr {
		return nil
	}
	k := pos[pre[pl]] - il // 中序遍历中的根结点位置
	root := &TreeNode{Val: pre[pl], Left: nil, Right: nil}
	root.Left = dfs(pre, in, pl+1, pl+k, il, il+k-1)
	root.Right = dfs(pre, in, pl+k+1, pr, il+k+1, ir)
	return root
}

func buildTree2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	var index int
	for i := range inorder {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}

	root.Left = buildTree2(preorder[1:index+1], inorder[:index])
	root.Right = buildTree2(preorder[index+1:], inorder[index+1:])

	return root
}
