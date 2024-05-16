package main

/*
	给你一个下标从 0 开始的整数数组 tasks ，其中 tasks[i] 表示任务的难度级别。在每一轮中，你可以完成 2 个或者 3 个 相同难度级别 的任务。

	返回完成所有任务需要的 最少 轮数，如果无法完成所有任务，返回 -1 。
*/

/*
	每轮完成的都是相同难度级别的任务，假设难度为 1 的任务有 c 个，问题变成：
	· 每轮可以把 c 减少 2，或者减少 3。把 c 减少到 0 最少要多少轮？
	例如 c = 10 时，3+3+2+2=10，最少要 4 轮。
	贪心地想，尽量多地「减少 3」，可以让轮数尽量少。

	分类讨论：
	· 如果 c = 1，无法完成，返回 -1。
	· 如果 c = 3k(k>=1)，只用「减少 3」就能完成，轮数为 c/3。
	· 如果 c = 3k+1(k>=1)，即 c = 3k'+4(k'>=0)，我们可以先把 c 减少到 4，然后使用两次「减少 2」，轮数为 (c-4)/3 + 2 = (c+2)/3 = ⌈c/3⌉
	· 如果 c = 3k+2(k>=1)，我们可以先把 c 减少到 2，然后使用一次「减少 2」，轮数为 (c-2)/3 + 1 = (c+1)/3 = ⌈c/3⌉。
	综上所述，对于 c(c>=2) 个相同难度任务级别的任务，最少需要操作 ⌈c/3⌉ = ⌊(c+2)/3⌋
*/

func minimumRounds(tasks []int) int {
	cnt := make(map[int]int)
	for _, t := range tasks {
		cnt[t]++
	}
	ans := 0
	for _, c := range cnt {
		if c == 1 {
			return -1
		}
		ans += (c + 2) / 3
	}
	return ans
}