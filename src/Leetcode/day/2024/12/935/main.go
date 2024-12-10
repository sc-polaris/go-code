package main

/*
	首先，由于无法从 5 移动到其他单元格，所以当 n≥2 时，马的初始位置不能等于 5。
	特判 n=1 的情况，直接返回 10。
	在下文中，n≥2。

	一、寻找子问题
	比如 n=3，我们要解决的问题（原问题）是：
	· 把马放在单元格 0 到 9（除了 5）上，然后移动 n−1=2 步，一共有多少种移动方案？
	枚举移动到的位置，比如一开始在 1，移动到 6，那么问题变成：
	· 把马放在单元格 6 上，然后移动 n−2=1 步，一共有多少种移动方案？
	这是和原问题相似的、规模更小的子问题，可以用递归解决。

	二、状态定义与状态转移方程（优化前）
	根据上面的讨论，我们需要在递归过程中跟踪以下信息：
	· i：还需要移动 i 步。
	· j：马在单元格 j 上。
	因此，定义状态为 dfs(i,j)，表示把马放在单元格 j 上，然后移动 i 步，有多少种移动方案。
	枚举马能移动到的单元格 k，问题变成：把马放在单元格 k 上，然后移动 i−1 步，有多少种移动方案，即 dfs(i−1,k)。
	累加得
				dfs(i,j) =  Σ dfs(i-1,k)
							k
	递归边界：dfs(0,j)=1。无法移动，算作一种移动方案，对应着电话号码 j。
			 9
	递归入口： ∑ dfs(n−1,j)，这是原问题，也是答案。注意 dfs(n−1,5)=0。
			j=0

*/

func knightDialer(n int) (ans int) {
	const mod = 1_000_000_007
	var next = [10][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}}
	var memo [5000][10]int
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p > 0 {
			return *p
		}
		res := 0
		for _, k := range next[j] {
			res += dfs(i-1, k)
		}
		res %= mod
		*p = res
		return res
	}

	if n == 1 {
		return 10
	}
	for j := range 10 {
		ans += dfs(n-1, j)
	}
	return ans % mod
}

/*
	根据对称性把这9个数分为 4 类：
	A：1 3 7 9
	B：2 8
	C：4 6
	D：0


	A B A
	C   C
	A B A
	  D

	还需要走 i 步		还需要走 i-1 步
			A	=	B' + C'
			B	=	2A'
			C	=	2A' + D'
			D	=	2C'
	令 j = 0,1,2,3 对应上面的 A，B，C，D类，那么有状态方程
			dfs(i,0) = dfs(i-1,1)+dfs(i-1,2)
			dfs(i,1) = 2*dfs(i-1,0)
			dfs(i,2) = 2*dfs(i-1,0)+dfs(i-1,3)
			dfs(i,3) = 2*dfs(i-1,2)
	答案为	4A+2B+2C+D
			4*dfs(n-1,0)+2*dfs(n-1,1)+2*dfs(n-1,2)+dfs(n-1,3)
*/

func knightDialer2(n int) (ans int) {
	const mod = 1_000_000_007
	var memo [5000][10]int
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i == 0 {
			return 1
		}
		p := &memo[i][j]
		if *p > 0 {
			return *p
		}
		if j == 0 {
			*p = (dfs(i-1, 1) + dfs(i-1, 2)) % mod
		} else if j == 1 {
			*p = 2 * dfs(i-1, 0) % mod
		} else if j == 2 {
			*p = (2*dfs(i-1, 0) + dfs(i-1, 3)) % mod
		} else {
			*p = 2 * dfs(i-1, 2) % mod
		}
		return *p
	}
	if n == 1 {
		return 10
	}
	return (4*dfs(n-1, 0) + 2*dfs(n-1, 1) + 2*dfs(n-1, 2) + dfs(n-1, 3)) % mod
}
