package main

/*
	给你一个整数 n 和一个二维整数数组 queries。

	有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。

	queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。

	所有查询中不会存在两个查询都满足 queries[i][0] < queries[j][0] < queries[i][1] < queries[j][1]。

	返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。
*/

/*
	方法一：区间并查集
	由于题目保证添加的边（捷径）不会交叉，从贪心的角度看，遇到捷径就走捷径是最优的。

	把目光放在边上。

	初始有 n−1 条边，我们在 0→1 这条边上，目标是到达 (n−2)→(n−1) 这条边，并把这条边走完。

	处理 queries 之前，需要走 n−1 条边。
	连一条从 2 到 4 的边，意味着什么？

	相当于把 2→3 这条边和 3→4 这条边合并成一条边。现在从起点到终点需要 3 条边。
	连一条从 0 到 2 的边，意味着什么？

	相当于把 0→1 这条边和 1→2 这条边合并成一条边。现在从起点到终点需要 2 条边。

	用并查集实现边的合并。初始化一个大小为 n−1 的并查集，并查集中的节点 i 表示题目的边 i→(i+1)。（相当于给每条边编号 0,1,2,…n−2。）

	连一条从 L 到 R 的边，相当于把并查集中的节点 L,L+1,L+2⋯,R−2 合并到并查集中的节点 R−1 上。

	合并的同时，维护并查集连通块个数。

	答案就是每次合并后的并查集连通块个数。
*/

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	fa := make([]int, n-1)
	for i := range fa {
		fa[i] = i
	}
	// 非递归并查集
	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	ans := make([]int, len(queries))
	cnt := n - 1 // 并查集连通块的个数
	for qi, q := range queries {
		l, r := q[0], q[1]-1
		fr := find(r)
		for i := find(l); i < r; i = find(i + 1) {
			fa[i] = fr
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}

/*
	方法二：记录跳转位置
	定义 nxt[i] 表示 i 指向的最右节点编号，这里 0≤i≤n−2。
	初始值 nxt[i]=i+1。

	连一条从 L 到 R 的边，分类讨论：
	· 如果之前连了一条从 L' 到 R′ 的边，且区间 [L,R] 被 [L′,R′] 包含，则什么也不做。
	· 否则更新 nxt[L]=R，在更新前，标记 [nxt[L],R−1] 中的没有被标记的点，表示这些点被更大的区间包含。怎么标记？把 nxt[i] 置为 r，这样可以把进入循环和继续循环的逻辑合并成一个：当 nxt[i]<r 时进入循环/继续循环。
	和方法一一样，维护一个 cnt 变量，每把一个 nxt[i] 置为 0，就把 cnt 减一。
*/

func shortestDistanceAfterQueries2(n int, queries [][]int) []int {
	nxt := make([]int, n-1)
	for i := range nxt {
		nxt[i] = i + 1
	}

	ans := make([]int, len(queries))
	cnt := n - 1
	for qi, q := range queries {
		for i, r := q[0], q[1]; nxt[i] < r; i, nxt[i] = nxt[i], r {
			cnt--
		}
		ans[qi] = cnt
	}
	return ans
}
