package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	isOdd := 0
	for len(q) > 0 {
		if isOdd != 0 {
			// 反转
			n := len(q)
			for i := 0; i < n/2; i++ {
				x, y := q[i], q[n-1-i]
				x.Val, y.Val = y.Val, x.Val
			}
		}
		tmp := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			if node.Left != nil {
				tmp = append(tmp, node.Left, node.Right)
			}
		}
		q = tmp
		isOdd ^= 1
	}
	return root
}

func main() {

}
