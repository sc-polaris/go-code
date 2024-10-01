package main

/*
	本题如果只求以 0 为根时的猜对次数 cnt0，把 guesses 转成哈希表，dfs 一次这棵树就可以算出来
	如果枚举 0 到 n-1 的每个点作为树根，就需要 dfs n 次，需要 O(n^2) 的时间，优化：

	注意要，如果节点 x 和节点 y 相邻，那么从「以 x 为根的树」变成「以 y 为根的树」，
	就只有 x 和 y 的父子关系改变了，其余相邻的父子关系没有变化。所以只有 [x,y] 和
	[y,x] 这两个猜测的正确性变了，其余猜测的正确性不变。

	因此，在计算出 cnt0 之后，我们可以再次从 0 出发，dfs 这棵树。从节点 x 递归到节点 y 时：
	· 如果有猜测 [x,y]，那么猜对次数减一。
	· 如果有猜测 [y,x]，那么猜对次数加一。
	dfs 的同时，统计猜对次数 >= k 的节点个数，即为答案。

*/

func rootCount(edges [][]int, guesses [][]int, k int) (ans int) {
	g := make([][]int, len(edges)+1)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	type pair struct{ x, y int }
	s := make(map[pair]int, len(guesses))
	for _, p := range guesses {
		s[pair{p[0], p[1]}] = 1
	}

	cnt0 := 0
	var dfs func(int, int)
	dfs = func(x, fa int) {
		for _, y := range g[x] {
			if y != fa {
				if s[pair{x, y}] == 1 { // 以 0 为根猜对了
					cnt0++
				}
				dfs(y, x)
			}
		}
	}
	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(x, fa, cnt int) {
		if cnt >= k { // 此时 cnt 就是以 x 为根时的猜对次数
			ans++
		}
		for _, y := range g[x] {
			if y != fa {
				reroot(y, x, cnt-s[pair{x, y}]+s[pair{y, x}])
			}
		}
	}
	reroot(0, -1, cnt0)
	return
}

func main() {

}
