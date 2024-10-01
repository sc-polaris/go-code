package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements map[int]bool

func Constructor(root *TreeNode) FindElements {
	f := FindElements{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, val int) {
		if node == nil {
			return
		}
		f[val] = true
		dfs(node.Left, val*2+1)
		dfs(node.Right, val*2+2)
	}
	dfs(root, 0)
	return f
}

func (f FindElements) Find(target int) bool {
	return f[target]
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

func main() {

}
