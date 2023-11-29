package dfs

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	var pre []*Node
	var dfs func(*Node, int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}
		if depth == len(pre) { // node是这一层最左边
			pre = append(pre, node)
		} else { // pre[depth] 是 node 左边的节点
			pre[depth].Next = node // node 左边的节点指向 node
			pre[depth] = node
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return root
}
