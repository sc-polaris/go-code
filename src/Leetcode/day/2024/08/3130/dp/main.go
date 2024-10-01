package main

/*
	和 dfs(i,j,k) 一样，定义 f[i][j][k] 表示用 i 个 0 和 j 个 1 构造稳定数组的方案数，其中第 i+j 个位置要填 k，其中 k 为 0 或 1。
	状态转移方程：
						f[i][j][0]=f[i−1][j][0]+f[i−1][j][1]−f[i−limit−1][j][1]
						f[i][j][1]=f[i][j−1][0]+f[i][j−1][1]−f[i][j−limit−1][0]

	如果 i≤limit 则 f[i−limit−1][j][1] 视作 0，
	如果 j≤limit 则 f[i][j−limit−1][0] 视作 0。

	初始值：	f[i][0][0]=f[0][j][1]=1，其中 1≤i≤min(limit,zero), 1≤j≤min(limit,one)。翻译自递归边界。
	答案：	f[zero][one][0]+f[zero][one][1]。翻译自递归入口。

*/

func numberOfStableArrays(zero int, one int, limit int) int {
	const mod = 1_000_000_007
	f := make([][][2]int, zero+1)
	for i := range f {
		f[i] = make([][2]int, one+1)
	}
	for i := 1; i <= min(limit, zero); i++ {
		f[i][0][0] = 1
	}
	for j := 1; j <= min(limit, one); j++ {
		f[0][j][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			f[i][j][0] = (f[i-1][j][0] + f[i-1][j][1]) % mod
			if i > limit {
				// + mod 保证答案非负
				f[i][j][0] = (f[i][j][0] - f[i-limit-1][j][1] + mod) % mod
			}
			f[i][j][1] = (f[i][j-1][0] + f[i][j-1][1]) % mod
			if j > limit {
				f[i][j][1] = (f[i][j][1] - f[i][j-limit-1][0] + mod) % mod
			}
		}
	}
	return (f[zero][one][0] + f[zero][one][1]) % mod
}
