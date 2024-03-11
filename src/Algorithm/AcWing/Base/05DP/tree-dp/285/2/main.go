package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 6010

type Node struct {
	children []*Node
	happy    int
	idx      int
}

var (
	in    = bufio.NewReader(os.Stdin)
	ot    = bufio.NewWriter(os.Stdout)
	n     int
	f     [N][2]int // f[u][0]: 所有以u为根的子树中选择，并且不选u的方案  f[u][1]: 所有以u为根的子树中选择，并且选u的方案
	hasFa [N]bool
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root *Node) {
	f[root.idx][1] = root.happy // 如果选当前节点
	for _, child := range root.children {
		dfs(child)
		f[root.idx][0] += max(f[child.idx][0], f[child.idx][1])
		f[root.idx][1] += f[child.idx][0]
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	nodes := make([]*Node, n+1)
	var h int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &h)
		nodes[i] = &Node{happy: h, idx: i}
	}

	for i := 1; i < n; i++ {
		var l, k int
		fmt.Fscan(in, &l, &k)
		hasFa[l] = true // 说明l他有爸爸（划掉）上司
		nodes[k].children = append(nodes[k].children, nodes[l])
	}

	root := 1 // 找根节点
	for hasFa[root] {
		root++
	}

	dfs(nodes[root])

	fmt.Fprintln(ot, max(f[root][0], f[root][1])) // 输出不选根节点与选根节点的最大值

}
