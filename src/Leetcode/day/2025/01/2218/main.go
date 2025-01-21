package main

/*
	一张桌子上总共有 n 个硬币 栈 。每个栈有 正整数 个带面值的硬币。

	每一次操作中，你可以从任意一个栈的 顶部 取出 1 个硬币，从栈中移除它，并放入你的钱包里。

	给你一个列表 piles ，其中 piles[i] 是一个整数数组，分别表示第 i 个栈里 从顶到底 的硬币面值。
	同时给你一个正整数 k ，请你返回在 恰好 进行 k 次操作的前提下，你钱包里硬币面值之和 最大为多少 。
*/

/*
	记忆化搜索：
	类似 0-1 背包，定义 dfs(i,j) 表示从 piles[0] 到 piles[i] 中，选体积之和至多为 j 的物品时，物品价值之和的最大值。
	枚举第 i 组的所有物品（枚举前缀和），设当前物品体积为 w，价值为 v，那么问题变成从前 i−1 个物品组中，选体积之和至多为
	j−w 的物品时，物品价值之和的最大值，即 dfs(i−1,j−w)，加上 v 得到 dfs(i,j)。
	所有情况取最大值，得
						dfs(i,j)=max dfs(i−1,j−w)+v
								(v,w)
	如果该组不选物品，则上式中的 v=w=0。
	递归边界：dfs(−1,j)=0。
	递归入口：dfs(n−1,k)，这是原问题，也是答案。
*/

func maxValueOfCoins(piles [][]int, k int) int {
	n := len(piles)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
	}
	var dfs func(int, int) int
	dfs = func(i, j int) (res int) {
		if i < 0 {
			return
		}
		p := &memo[i][j]
		if *p != 0 {
			return *p
		}
		defer func() { *p = res }()

		// 不选这一组中的任何物品
		res = dfs(i-1, j)
		// 枚举选哪个
		v := 0
		for w := range min(j, len(piles[i])) {
			v += piles[i][w]
			// w 从 0 开始，物品体积为 w + 1
			res = max(res, dfs(i-1, j-w-1)+v)
		}
		return
	}
	return dfs(n-1, k)
}

/*
	递推

	具体来说，f[i+1][j] 的定义和 dfs(i,j) 的定义是一样的，都表示从 piles[0] 到 piles[i] 中，选体积之和至多为 j 的物品时，
	物品价值之和的最大值。这里 +1 是为了把 dfs(−1,j) 这个状态也翻译过来，这样我们可以把 f[0][j] 作为初始值。

*/

func maxValueOfCoins2(piles [][]int, k int) int {
	f := make([][]int, len(piles)+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	for i, pile := range piles {
		for j := range k + 1 {
			// 不选这一组
			f[i+1][j] = f[i][j]
			// 枚举选哪个
			v := 0
			for w := range min(j, len(pile)) {
				v += pile[w]
				// w 从 0 开始，物品体积为 w+1
				f[i+1][j] = max(f[i+1][j], f[i][j-w-1]+v)
			}
		}
	}
	return f[len(piles)][k]
}

/*
	空间优化
*/

func maxValueOfCoins3(piles [][]int, k int) int {
	f := make([]int, k+1)
	sumN := 0
	for _, pile := range piles {
		n := len(pile)
		for i := 1; i < n; i++ {
			pile[i] += pile[i-1] // 提前计算 pile 的前缀和
		}
		sumN = min(sumN+n, k)
		for j := sumN; j > 0; j-- { // 优化：j 从前 i 个栈的大小之和开始枚举
			for w, v := range pile[:min(n, j)] {
				f[j] = max(f[j], f[j-w-1]+v) // w 从 0 开始，物品体积为 w+1
			}
		}
	}
	return f[k]
}
