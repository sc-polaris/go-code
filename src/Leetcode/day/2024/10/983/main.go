package main

/*
	在一个火车旅行很受欢迎的国度，你提前一年计划了一些火车旅行。在接下来的一年里，你要旅行的日子将以一个名为 days 的数组给出。每一项是一个从 1 到 365 的整数。

	火车票有 三种不同的销售方式 ：
	· 一张 为期一天 的通行证售价为 costs[0] 美元；
	· 一张 为期七天 的通行证售价为 costs[1] 美元；
	· 一张 为期三十天 的通行证售价为 costs[2] 美元。
	通行证允许数天无限制的旅行。 例如，如果我们在第 2 天获得一张 为期 7 天 的通行证，那么我们可以连着旅行 7 天：第 2 天、第 3 天、第 4 天、第 5 天、第 6 天、第 7 天和第 8 天。

	返回 你想要完成在给定的列表 days 中列出的每一天的旅行所需要的最低消费 。
*/

/*
	定义 dfs(i) 表示 1 到 i 天的最小花费。
*/

// 记忆化搜索
func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	isTravel := make([]bool, lastDay+1)
	for _, day := range days {
		isTravel[day] = true
	}
	memo := make([]int, lastDay+1)
	var dfs func(int) int
	dfs = func(i int) (res int) {
		if i <= 0 {
			return
		}
		p := &memo[i]
		if *p > 0 {
			return *p
		}
		defer func() { *p = res }()
		if !isTravel[i] { // 如果第 i 天不在 days 中，那么问题变成 1 到 i−1 天的最小花费，即
			return dfs(i - 1)
		}
		return min(dfs(i-1)+costs[0], dfs(i-7)+costs[1], dfs(i-30)+costs[2])
	}
	return dfs(lastDay)
}

// 递推
func mincostTickets2(days []int, costs []int) int {
	lastDay := days[len(days)-1]
	isTravel := make([]bool, lastDay+1)
	for _, day := range days {
		isTravel[day] = true
	}
	f := make([]int, lastDay+1)
	for i := 1; i <= lastDay; i++ {
		if !isTravel[i] {
			f[i] = f[i-1]
		} else {
			f[i] = min(f[i-1]+costs[0], f[max(i-7, 0)]+costs[1], f[max(i-30, 0)]+costs[2])
		}
	}
	return f[lastDay]
}

// 如果把数据范围修改为 days[i]≤10^9 ，上面的做法就不行了。能不能做到时间复杂度和 D 无关呢？比如只和 days 的长度 n 有关？
// 三指针优化
func mincostTickets3(days []int, costs []int) int {
	n := len(days)
	f := make([]int, n+1)
	j, k := 0, 0
	for i, d := range days {
		for days[j] <= d-7 {
			j++
		}
		for days[k] <= d-30 {
			k++
		}
		f[i+1] = min(f[i]+costs[0], f[j]+costs[1], f[k]+costs[2])
	}
	return f[n]
}
