package main

/*
	空间优化：

	观察上面的状态转移方程，在 计算 f[i] 时，不会用到 下标小于 i-1 的状态。

	所以可以去掉第一个唯独，反复利用同一个长为 n 的一维数组。

	为避免状态被覆盖，可以用一个变量 pre 记录 f[i-1][j]。
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
	f := make([]int, n)
	for i := 0; ; i++ {
		pre := 0
		for j, d := range dist[:n-1] {
			tmp := f[j+1]
			f[j+1] = (f[j] + d + speed - 1) / speed * speed
			if i > 0 {
				f[j+1] = min(f[j+1], pre+d)
			}
			pre = tmp
		}
		if f[n-1]+dist[n-1] <= speed*hoursBefore {
			return i
		}
	}
}
