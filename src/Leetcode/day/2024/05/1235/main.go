package main

import (
	"slices"
	"sort"
)

/*
	你打算利用空闲时间来做兼职工作赚些零花钱。
	这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。
	给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。
	注意，时间上出现重叠的 2 份工作不能同时进行。
	如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。
*/

/*
	分类讨论，求出按照结束时间排序后的前 i 个工作的最大报酬：
	1. 不选第 i 个工作，那么最大报酬等于前 i-1 个工作的最大报酬（转换成了一个规模更小的子问题）。
	2. 选第 i 个工作，由于时间不能重叠，设 j 是最大的满足 endTime[j] <= startTime[i] 的 j，那么
	   最大报酬等于前 j 个工作的最大报酬加上 profit[i]（同样转换成了一个规模更小的子问题）。
	3. 这两种决策取最大值。
	注意：由于按照结束时间排序，前 j 个工作中任意一个都不会与第 i 个工作的时间重叠。

	实现：
	定义 f[i] 表示按照结束时间排序后的前 i 个工作的最大报酬，用「选或不选」分类讨论：
	1. 不选第 i 个工作：f[i] = f[i-1]；
	2. 选第 i 个工作：f[i] = f[j] + profit[i]，其中 j 是最大的满足 endTime[j] <= startTime[i] 的 j，不存在时为 -1。
	两者取最大值，即
			f[i] = max(f[i-1],f[j]+profit[i])
	由于 i = 0 时会产生 -1，可以在 f 数组前面插入一个 0，与 f 有关的下标都 +1，即
			f[i+1] = max(f[i],f[j+1]+profit[i])
	初始项 f[0] = 0，答案为 f[n]。
*/

/*
	标准库二分的灵活运用
	1. >=	在有序数组中查询大于或等于某个数的最小数；
	2. >	在有序数组中查询大于某个数的最小数；
	3. <=	在有序数组中查询小于或等于某个数的最大数；
	4. <	在有序数组中查询小于某个数的最大数。

	一般编程语言的标准库中的二分只提供了查询 >= 和 > 的功能，并没有提供查询 <= 和 < 的功能。

	没有关系，稍微转换下就能解决。比如查询 > 得到了下标 i，那么 i-1 就是 <= 的结果了（假设数组为升序），同理 < 可以用 >= 算出来。
	注：> 和 >= 也可以转换，对于整数来说，> x 等价于 >= x+1
*/

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	type job struct{ start, end, profit int }
	jobs := make([]job, n)
	for i, start := range startTime {
		jobs[i] = job{start, endTime[i], profit[i]}
	}
	slices.SortFunc(jobs, func(a, b job) int { return a.end - b.end }) // 按照结束时间排序

	f := make([]int, n+1)
	for i, job := range jobs {
		j := sort.Search(i, func(j int) bool { return jobs[j].end > job.start })
		// 状态转移中，为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还是要 +1，抵消了
		f[i+1] = max(f[i], f[j]+job.profit)
	}
	return f[n]
}
