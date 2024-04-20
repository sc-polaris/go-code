package dfs

/*
	题干：
	给你一个整数 hoursBefore ，表示你要前往会议所剩下的可用小时数。要想成功抵达会议现场，你必须途经 n 条道路。道路的长度用一个长度为 n 的整数数组 dist 表示，
	其中 dist[i] 表示第 i 条道路的长度（单位：“千米”）。另给你一个整数 speed ，表示你在道路上前进的速度（单位：“千米每小时”）。

	当你通过第 i 条路之后，就必须休息并等待，直到 “下一个整数小时” 才能开始继续通过下一条道路。注意：你不需要在通过最后一条道路后休息，因为那时你已经抵达会议现场。

	· 例如，如果你通过一条道路用去 1.4 小时，那你必须停下来等待，到 2 小时才可以继续通过下一条道路。如果通过一条道路恰好用去 2 小时，就无需等待，可以直接继续。
	  然而，为了能准时到达，你可以选择 “跳过” 一些路的休息时间，这意味着你不必等待下一个整数小时。注意，这意味着与不跳过任何休息时间相比，你可能在不同时刻到达接
	  下来的道路。
	· 例如，假设通过第 1 条道路用去 1.4 小时，且通过第 2 条道路用去 0.6 小时。跳过第 1 条道路的休息时间意味着你将会在恰好 2 小时完成通过第 2 条道路，且你能
	  够立即开始通过第 3 条道路。

	返回准时抵达会议现场所需要的 “最小跳过次数” ，如果 “无法准时参会” ，返回 -1 。
*/

/*
	一、寻找子问题
	考虑枚举最小跳过次数，由于 dist[n-1] 无需跳过，我们从 dist[n-2] 开始思考。
	以示例 2 的 dist=[7,3,5,5] 为例，如果最多跳 2 次，用 「选或不选」分类讨论：
	· 不跳过 dist[n-2]=5：问题变成在最多跳过 2 次的情况下，经过 [7,3] 需要的最小时间。
	· 跳过	dist[n-2]=5：问题变成在最多跳过 1 次的情况下，经过 [7,3] 需要的最小时间。

	由于是否跳过都会把原问题变成一个 “和原问题相似的、规模更小的子问题” ，所以可以用 “递归” 解决。
	注：动态规划有「选或不选」和「枚举选哪个」两种基本思考方式。在做题时，可根据题目要求，选择适合题目的一种来思考。本题用到的是「选或不选」。

	二、状态定义与状态转移
	因为要解决的问题都形如「在最多跳过 i 次的情况下，从 dist[0] 到 dist[j] 需要的最小时间」，所以把它定义成 dfs(i,j)。请注意，如果在 j
	处选择不跳过的话，dfs(i,j) 需包含在 j 处等待的时间。

	考虑 dist[j] 是否跳过：
	· 不跳过：先算出在最多跳过 i 次的情况下，从 dist[0] 到 dist[j-1] 需要的最小时间，即 dfs(i,j-1)，然后加上 dist[j] 需要的时间，再
	  休息并等待（上取整），得
									 ⌈	  		    dist[j]	⌉
							dfs(i,j)=| dfs(i,j-1) + ------- |
									 |               speed  |
dfs(i,j) =｜ dfs(i,j-1) + dist[j]/speed   --- 上取整
	· 跳过：先算出在最多跳过 i-1 次的情况下，从 dist[0] 到 dist[j-1] 需要的最小时间，即 dfs(i-1,j-1)，然后加上 dist[j] 需要的时间，
	  得：
													  dist[j]
							dfs(i,j) = dfs(i-1,j-1) + -------
													   speed
	这两种情况取最小值，就得到了 dfs(i,j)，即
											⌈	  		   dist[j] ⌉                  dist[j]
							dfs(i,j) = min	| dfs(i,j-1) + ------- | , dfs(i-1,j-1) + -------
											|               speed  |                   speed

	递归边界：dfs(i,-1) = 0。
	递归入口：dfs(i,n-2)。注意 dist[n-1] 无需跳过，单独计算。
	最小的满足
			dfs(i,n-2) + dist[n-1]/speed <= hoursBefore
	的 i 就是答案

	三、改进：避免浮点误差
	注意上面的计算过程包含除法，所以 dfs(i,j) 不一定是个整数。但如果使用浮点数存储 dfs(i,j) 的话，会因为 “舍入误差” 导致计算结果可能与
	实际不符。为避免浮点运算带来的误差，可以把 dfs(i,j) 的定义改成在同等时间下用 speed 速度能走完的 “距离”。换句话说，在最多跳过 i 次的
	情况下，从 dist[0] 到 dist[j] 需要的最小时间是 dfs(i,j)/speed。

	考虑 dist[j] 是否跳过：
					⌈ dfs(i,j-1)	dist[j]	⌉					⌈ dfs(i,j-1) + dist[j] ⌉
	· 不跳过：用时为	| ----------  + ------- |，所以有 dfs(i,j) =	| -------------------- |
					|    speed       speed  |					|         speed        |

					dfs(i-1,j-1)	dist[j] 					dfs(i,j-1) + dist[j]
	· 跳过：用时为	------------  + ------- ，所以有 dfs(i,j) =	-------------------- * speed = dfs(i-1,j-1) + dist[j]
					   speed         speed  					        speed

	二者取最小值，状态转移变成
											⌈ dfs(i,j-1) + dist[j] ⌉
							dfs(i,j) = min	| -------------------- |, dfs(i-1,j-1) + dist[j]
											|         speed        |
	最小的满足
			dfs(i,n-2)/speed + dist[n-1]/speed <= hourBefore
	即
			dfs(i,n-2) + dist[n-1] <= speed * hourBefore
	的 i 就是答案

	这样就完全不设计浮点运算了！
	在递归钱可以先判断是否满都 sum(dist[0:n-1]) <= speed*hourBefore 如果不满足则返回 -1

	上取整：
							⌈ a ⌉	|a + b - 1|
							|---| = |---------|
							| b |   ⌊    b    ⌋
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
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j < 0 { // 递归边界
			return 0
		}
		p := &memo[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := (dfs(i, j-1) + dist[j] + speed - 1) / speed * speed
		if i > 0 {
			res = min(res, dfs(i-1, j-1)+dist[j])
		}
		*p = res // 记忆化
		return res
	}
	for i := 0; ; i++ {
		if dfs(i, n-2)+dist[n-1] <= speed*hoursBefore {
			return i
		}
	}
}
