package main

import "math"

/*
	给你两个长度为 n 下标从 0 开始的整数数组 cost 和 time ，分别表示给 n 堵不同的墙刷油漆需要的开销和时间。你有两名油漆匠：
	· 一位需要 付费 的油漆匠，刷第 i 堵墙需要花费 time[i] 单位的时间，开销为 cost[i] 单位的钱。
	· 一位 免费 的油漆匠，刷 任意 一堵墙的时间为 1 单位，开销为 0 。但是必须在付费油漆匠 工作 时，免费油漆匠才会工作。
	请你返回刷完 n 堵墙最少开销为多少。
*/

/*
	方法一：付费与免费的时间差
	用「选或不选」的来思考，即是否付费刷墙。
	考虑第 n−1 堵墙是否付费刷：
	· 选择付费刷第 n−1 堵墙，那么问题变成：刷前 n−2 堵墙，在「付费时间之和为 time[n−1]，免费时间之和为 0」状态下的最少开销。
	· 选择不付费，即免费刷第 n−1 堵墙，那么问题变成：刷前 n−2 堵墙，在「付费时间之和为0，免费时间之和为 1」状态下的最少开销。
	定义：dfs(i,j,k) 表示刷前 i 堵墙，在「付费时间之和为 j，免费时间之和为 k」状态下的最少开销。
	递归到终点时，如果 j ≥ k，说明这种方案是合法的，否则不合法。
	但这样定义的话，状态个数就太多了，需要优化。
	由于最后是比较的 j 和 k 的「相对大小」，那么不妨把 j 重新定义为「付费时间之和」减去「免费时间之和」，这样递归到终点时，如果
	j ≥ 0，说明这种方案是合法的，否则不合法。这样一来，状态个数就大大减少了。

	分类讨论：
	· 选择付费刷第 i 堵墙：dfs(i,j) = dfs(i−1,j+time[i])+cost[i]。
	· 选择免费刷第 i 堵墙：dfs(i,j) = dfs(i−1,j−1)。
	两种情况取最小值，得
					dfs(i,j) = min(dfs(i−1,j+time[i])+cost[i],dfs(i−1,j−1))
	递归边界：如果 j > i，那么剩余的墙都可以免费刷，即 dfs(i,j) = 0，否则 dfs(−1,j) = ∞。
	递归入口：dfs(n−1,0)，即答案。

*/

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j > i { // 剩余的墙都可以免费刷
			return 0
		}
		if i < 0 { // 上面 if 不成立，意味着 j < 0，不符合题目要求
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j+n] // 加上偏移量 n，防止出现负数
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, 0)
}

/*
	01 背包
	根据题意，付费刷墙个数 + 免费刷墙个数 = n。
	同时，付费刷墙时间之和必须 ≥ 免费刷墙个数。
	结合这两个式子，得到：付费刷墙时间之和 ≥ n− 付费刷墙个数。
	移项，得到：「付费刷墙时间+1」之和 ≥ n。（把个数拆分成 1+1+1+⋯）
	把 time[i]+1 看成物品体积，cost[i] 看成物品价值，问题变成：
	· 从 n 个物品中选择体积和至少为 n 的物品，价值和最小是多少？
	这是 0-1 背包的一种「至少装满」的变形。我们可以定义 dfs(i,j) 表示考虑前 i 个物品，剩余还需要凑出 j 的体积，此时的最小价值和。
	和 0-1 背包一样，用选或不选思考，可以得到类似的状态转移方程：
		dfs(i,j) = min(dfs(i−1,j−time[i]−1)+cost[i],dfs(i−1,j))
	递归边界：如果 j ≤ 0，不需要再选任何物品了，返回 0；如果 i < 0，返回无穷大。
	递归入口：dfs(n−1,n)，表示体积和至少为 n，这正是我们要计算的。
*/

func paintWalls2(cost []int, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1 // 没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int { // j 表示剩余需要的体积
		if j <= 0 { // 没有约束，后面什么也不用选了
			return 0
		}
		if i < 0 { // 此时 j > 0，但没有物品可选，不合法
			return math.MaxInt / 2 // 防止加法溢出
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		*p = min(dfs(i-1, j-time[i]-1)+cost[i], dfs(i-1, j))
		return *p
	}
	return dfs(n-1, n)
}

/*
	01 背包 递推
*/

func paintWalls3(cost []int, time []int) int {
	n := len(cost)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt / 2 // 防止加法溢出
	}
	for i, c := range cost {
		t := time[i] + 1
		for j := n; j >= 0; j-- {
			f[j] = min(f[j], f[max(j-t, 0)]+c)
		}
	}
	return f[n]
}
