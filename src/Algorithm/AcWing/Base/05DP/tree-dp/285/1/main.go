package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 6010

var (
	in       = bufio.NewReader(os.Stdin)
	ot       = bufio.NewWriter(os.Stdout)
	n        int
	h, e, ne = make([]int, N), make([]int, N), make([]int, N)
	idx      int
	happy    [N]int
	f        [N][2]int // f[u][0]: 所有以u为根的子树中选择，并且不选u的方案  f[u][1]: 所有以u为根的子树中选择，并且选u的方案
	hasFa    [N]bool
)

func add(a, b int) { e[idx] = b; ne[idx] = h[a]; h[a] = idx; idx++ }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func memset(a []int, v int) {
	for i := range a {
		a[i] = v
	}
}

func dfs(u int) {
	f[u][1] = happy[u] // 如果选当前节点
	for i := h[u]; i != -1; i = ne[i] {
		j := e[i]
		dfs(j) // 回溯
		f[u][0] += max(f[j][1], f[j][0])
		f[u][1] += f[j][0]
	}
}

func main() {
	defer ot.Flush()

	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &happy[i])
	}

	memset(h, -1)

	for i := 1; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		hasFa[a] = true // 说明a他有爸爸（划掉）上司
		add(b, a)       // 把a加入到b的后面, b是上司
	}

	root := 1 // 找根节点
	for hasFa[root] {
		root++
	}

	dfs(root)

	fmt.Fprintln(ot, max(f[root][0], f[root][1])) // 输出不选根节点与选根节点的最大值

}
