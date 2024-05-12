package dfs

/*
	厨房里总共有 n 个橘子，你决定每一天选择如下方式之一吃这些橘子：

	· 吃掉一个橘子。
	· 如果剩余橘子数 n 能被 2 整除，那么你可以吃掉 n/2 个橘子。
	· 如果剩余橘子数 n 能被 3 整除，那么你可以吃掉 2*(n/3) 个橘子。
	每天你只能从以上 3 种方案中选择一种方案。

	请你返回吃掉所有 n 个橘子的最少天数。
*/

func minDays(n int) int {
	memo := make(map[int]int)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return i
		}
		if v, ok := memo[i]; ok { // 之前计算过
			return v
		}
		res := min(dfs(i/2)+i%2, dfs(i/3)+i%3) + 1
		memo[i] = res // 记忆化
		return res
	}
	return dfs(n)
}
