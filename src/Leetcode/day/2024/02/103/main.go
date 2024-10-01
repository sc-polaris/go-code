package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*TreeNode{root}
	for even := false; len(q) > 0; even = !even {
		n := len(q)
		values := make([]int, n) // 大小已知
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if even {
				values[n-1-i] = node.Val // 倒着添加
			} else {
				values[i] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, values)
	}
	return
}

func main() {

}
