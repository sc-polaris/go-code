package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, c := range node.Children {
			dfs(c)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}

func main() {

}
