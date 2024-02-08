package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	var fx, fy *TreeNode
	var dx, dy int
	var hasX, hasY bool

	// 用来判断是否遍历到 x 或 y 的辅助函数
	update := func(node, parent *TreeNode, depth int) {
		if node.Val == x {
			fx, dx, hasX = parent, depth, true
		} else if node.Val == y {
			fy, dy, hasY = parent, depth, true
		}
	}

	type pair struct {
		node  *TreeNode
		depth int
	}

	q := []pair{{root, 0}}
	update(root, nil, 0)
	for len(q) > 0 && (!hasX || !hasY) {
		node, depth := q[0].node, q[0].depth
		q = q[1:]
		if node.Left != nil {
			q = append(q, pair{node.Left, depth + 1})
			update(node.Left, node, depth+1)
		}
		if node.Right != nil {
			q = append(q, pair{node.Right, depth + 1})
			update(node.Right, node, depth+1)
		}
	}

	return dx == dy && fx != fy
}

func main() {

}
