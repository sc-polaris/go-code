package main

/*
	对于节点 x，我们需要计算节点值与 nums[x] 互质的最近祖先节点是哪个。

	最暴力的做法是，枚举 x 的所有祖先节点。但如果这棵树是一条链。枚举 x 的所有祖先节点需要
	O(n) 的时间，每个点都这样枚举的话，总共需要 O(n^2) 的时间，太慢了。

	注意到，所有节点的节点值不超过 50，我们可以枚举 [1,50] 中与 nums[x] 互质的数。由于要计
	算的是 「最近」祖先，对于节点值相同的祖先，只需要枚举深度最大的。因此，对于节点 x，我们至
	多枚举它的 50 个祖先。这样总共需要 O(nU) 的时间，其中 U=50。

	具体来说，我们需要在递归这棵树的同时，维护两组信息：
	1. valDepth 数组。其中 valDepth[j] 保存节点值等于 j 的最近祖先的深度。
	2. valNodeId 数组。其中 valNodeId[j] 保存节点值等于 j 的最近祖先的节点编号。
	设当前节点值为 val=nums[x]，我们枚举 [1,50] 中与 val 互质的数字 j，计算出 valDepth[j]
	的最大值，及其对应的节点编号，即为答案 ans[x]。
	代码实现时，可与你预处理 [1,50] 中有哪些数对是互质的。

*/

const mx = 51

var coprime [mx][]int

func init() {
	// 预处理：coprime[i] 保存 [1, MX) 中与 i 互质的所有元素
	for i := 1; i < mx; i++ {
		for j := 1; j < mx; j++ {
			if gcd(i, j) == 1 {
				coprime[i] = append(coprime[i], j)
			}
		}
	}
}

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	type pair struct{ depth, id int }
	valDepthId := [mx]pair{}
	var dfs func(int, int, int)
	dfs = func(x, fa, depth int) {
		val := nums[x] // x 的节点值
		// 计算与 val 互质的数中，深度最大的节点编号
		maxDepth := 0
		for _, j := range coprime[val] {
			p := valDepthId[j]
			if p.depth > maxDepth {
				maxDepth = p.depth
				ans[x] = p.id
			}
		}

		tmp := valDepthId[val]           // 用于恢复现场
		valDepthId[val] = pair{depth, x} // 保存 val 对应节点深度和节点编号
		for _, y := range g[x] {
			if y != fa {
				dfs(y, x, depth+1)
			}
		}
		valDepthId[val] = tmp // 恢复现场
	}
	dfs(0, -1, 1)
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
