package main

/*
	给你一个炸弹列表。一个炸弹的 爆炸范围 定义为以炸弹为圆心的一个圆。

	炸弹用一个下标从 0 开始的二维整数数组 bombs 表示，其中 bombs[i] = [xi, yi, ri] 。xi 和 yi 表示第 i 个炸弹的 X 和 Y 坐标，ri 表示爆炸范围的 半径 。

	你需要选择引爆 一个 炸弹。当这个炸弹被引爆时，所有 在它爆炸范围内的炸弹都会被引爆，这些炸弹会进一步将它们爆炸范围内的其他炸弹引爆。

	给你数组 bombs ，请你返回在引爆 一个 炸弹的前提下，最多 能引爆的炸弹数目。

	方法一：建图 + 枚举起点 DFS

	问：为什么不能用并查集？

	答：注意本题是有向图。例如炸弹 0 可以引爆炸弹 2，炸弹 1 可以引爆炸弹 2，对应有向边 0→2,1→2，那么正确答案是 2。如果用并查集做的话，
	会把 0,1,2 三个点合并起来，计算出错误的答案 3。
*/

func maximumDetonation(bombs [][]int) (ans int) {
	n := len(bombs)
	g := make([][]int, n)
	for i, p := range bombs {
		x, y, r := p[0], p[1], p[2]
		for j, q := range bombs {
			dx := x - q[0]
			dy := y - q[1]
			if j != i && dx*dx+dy*dy <= r*r {
				g[i] = append(g[i], j) // i 可以引爆 j
			}
		}
	}

	vis := make([]bool, n)
	var dfs func(int) int
	dfs = func(x int) int {
		vis[x] = true
		cnt := 1
		for _, y := range g[x] {
			if !vis[y] {
				cnt += dfs(y)
			}
		}
		return cnt
	}
	for i := range g {
		clear(vis)
		ans = max(ans, dfs(i))
	}
	return
}
