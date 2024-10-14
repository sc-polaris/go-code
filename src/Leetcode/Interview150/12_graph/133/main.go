package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	vis := make(map[*Node]*Node)
	var dfs func(*Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		// 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
		if _, ok := vis[node]; ok {
			return vis[node]
		}
		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
		cloneNode := &Node{node.Val, []*Node{}}
		vis[node] = cloneNode
		// 遍历该节点的邻居并更新克隆节点的邻居列表
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(n))
		}
		return cloneNode
	}
	return dfs(node)
}

func cloneGraph2(node *Node) *Node {
	if node == nil {
		return nil
	}
	vis := make(map[*Node]*Node)
	q := []*Node{node}
	vis[node] = &Node{node.Val, []*Node{}}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, ne := range n.Neighbors {
			if _, ok := vis[ne]; !ok {
				vis[ne] = &Node{ne.Val, []*Node{}}
				q = append(q, ne)
			}
			vis[n].Neighbors = append(vis[n].Neighbors, vis[ne])
		}
	}
	return vis[node]
}
