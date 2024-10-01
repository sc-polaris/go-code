package main

import "slices"

/*
	你有 n 个工作和 m 个工人。给定三个数组： difficulty, profit 和 worker ，其中:
	· difficulty[i] 表示第 i 个工作的难度，profit[i] 表示第 i 个工作的收益。
	· worker[i] 是第 i 个工人的能力，即该工人只能完成难度小于等于 worker[i] 的工作。

	每个工人 最多 只能安排 一个 工作，但是一个工作可以 完成多次 。
	· 举个例子，如果 3 个工人都尝试完成一份报酬为 $1 的同样工作，那么总收益为 $3 。如果一个工人不能完成任何工作，他的收益为 $0 。

	返回 在把工人分配到工作岗位后，我们所能获得的最大利润
*/

func maxProfitAssignment(difficulty []int, profit []int, worker []int) (ans int) {
	n := len(difficulty)
	type job struct{ d, p int }
	jobs := make([]job, n)
	for i, d := range difficulty {
		jobs[i] = job{d, profit[i]}
	}
	slices.SortFunc(jobs, func(a, b job) int { return a.d - b.d })
	slices.Sort(worker)
	j, maxProfit := 0, 0
	for _, w := range worker {
		for j < n && jobs[j].d <= w {
			maxProfit = max(maxProfit, jobs[j].p)
			j++
		}
		ans += maxProfit
	}
	return
}
