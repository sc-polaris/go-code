package main

/*
	翻译成递推
	f[i][j] 表示在最多跳过 i 次的情况下，从 dist[0] 到 dist[j] 需要的最小时间，再乘上 speed

	相应的递推式（状态转移方程）也和 dfs 一样
											⌈ f[i][[j-1] + dist[j] ⌉
							f[i][j] = min	| -------------------- |, f[i-1][j-1] + dist[j]
											|          speed       |
	但是，这种定义方式没有状态能表示递归边界，即 j = -1 的情况

	解决办法：在 二位数组 f 的左侧插入一列状态，那么其余状态全部向右移动一位，f 的下标 j 需要加一，也就是把 f[·][j] 改成 f[·][j+1]，
			把 f[·][j-1] 改为 f[·][j]。
	修改后 f[i][j+1] 表示在最多跳过 i 次的情况下，从 dist[0] 到 dist[j] 需要的最小时间，再乘上 speed。此时 f[0] 就对应递归边界了。
	修改后的递推式为
											⌈ f[i][[j] + dist[j] ⌉
							f[i][j+1] = min	| ------------------ | * speed, f[i-1][j] + dist[j]
											|        speed       |
	初始值 f[i][0] = 0，翻译自递归边界 dfs(i,-1) = 0。
	最小的满足
							f[i][n-1] + dist[n-1] <= speed * hourBefore
	的 i 就是答案。

*/

func minSkips(dist []int, speed, hoursBefore int) int {
	sumDist := 0
	for _, d := range dist {
		sumDist += d
	}
	if sumDist > speed*hoursBefore {
		return -1
	}

	n := len(dist)
	f := make([][]int, n)
	for i := 0; ; i++ {
		f[i] = make([]int, n)
		for j, d := range dist[:n-1] {
			f[i][j+1] = (f[i][j] + d + speed - 1) / speed * speed
			if i > 0 {
				f[i][j+1] = min(f[i][j+1], f[i-1][j]+d)
			}
		}
		if f[i][n-1]+dist[n-1] <= speed*hoursBefore {
			return i
		}
	}
}
