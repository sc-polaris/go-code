package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	cur := []*TreeNode{root}
	for len(cur) > 0 {
		var nxt []*TreeNode
		values := make([]int, len(cur)) // 大小已知
		for i, node := range cur {
			values[i] = node.Val
			if node.Left != nil {
				nxt = append(nxt, node.Left)
			}
			if node.Right != nil {
				nxt = append(nxt, node.Right)
			}
		}
		cur = nxt
		ans = append(ans, values)
	}
	return
}

func main() {

}
