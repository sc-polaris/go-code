package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printFromTopToBottom(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		res = append(res, t.Val)
		if t.Left != nil {
			q = append(q, t.Left)
		}
		if t.Right != nil {
			q = append(q, t.Right)
		}
	}

	return res
}
func main() {

}
