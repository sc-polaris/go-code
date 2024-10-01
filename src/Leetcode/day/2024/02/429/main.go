package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (ans [][]int) {
	if root == nil {
		return
	}
	q := []*Node{root}
	for len(q) > 0 {
		n := len(q)
		vals := make([]int, n)
		for i := range vals {
			node := q[0]
			q = q[1:]
			vals[i] = node.Val
			q = append(q, node.Children...)
		}
		ans = append(ans, vals)
	}
	return
}

func main() {

}
