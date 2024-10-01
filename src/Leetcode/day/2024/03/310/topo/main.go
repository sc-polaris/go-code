package main

/*

	1. 首先找到所有度为 1 的节点压入队列，此时令节点剩余计数 remainNodes = n
	2. 同时将当前 remainNodes 计数减去出度为 1 的节点数目，将最外层度为 1 的叶子节点取出，
		并将与之相邻的节点的度减少，重复上述步骤将当前节中度为 1 的节点压入队列
	3. 重复上述步骤，直到剩余节点数数组 remainNodes <= 2 时，此时剩余的节点即为当前高度最小树的根节点。

*/

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)
	degree := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		degree[x]++
		degree[y]++
	}
	var q []int
	for i, d := range degree {
		if d == 1 {
			q = append(q, i)
		}
	}
	remainNodes := n
	for remainNodes > 2 {
		remainNodes -= len(q)
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, y := range g[x] {
				degree[y]--
				if degree[y] == 1 {
					q = append(q, y)
				}
			}
		}
	}
	return q
}

func main() {

}
