package main

import "fmt"

func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	cnt := make([]int, n)
	for _, t := range trips {
		end := t[1]
		var dfs func(int, int) bool
		dfs = func(x int, fa int) bool {
			if x == end {
				cnt[x]++
				return true // 找到 end
			}
			for _, y := range g[x] {
				if y != fa && dfs(y, x) {
					cnt[x]++ // x 是 end 的祖先节点，也就在路径上
					return true
				}
			}
			return false
		}
		dfs(t[0], -1)
	}

	var dfs func(int, int) (int, int)
	dfs = func(x int, fa int) (int, int) {
		notHalve := price[x] * cnt[x] // x 不变
		halve := notHalve / 2         // x 减半
		for _, y := range g[x] {
			if y != fa {
				nh, h := dfs(y, x)     // 计算 y 不变/减半的最小价值总和
				notHalve += min(nh, h) // x 不变，那么 y 可以不变，可以减半，取这两种情况的最小值
				halve += nh            // x 减半，那么 y 只能不变
			}
		}
		return notHalve, halve
	}
	return min(dfs(0, -1))
}

func main() {
	fmt.Println(minimumTotalPrice(
		4,
		[][]int{{0, 1}, {1, 2}, {1, 3}},
		[]int{2, 2, 10, 6},
		[][]int{{0, 3}, {2, 1}, {2, 3}},
	))
}
