package main

/*
	给你一个下标从 0 开始的二维整数矩阵 grid，大小为 n * n ，其中的值在 [1, n^2] 范围内。除了 a 出现 两次，b 缺失 之外，每个整数都 恰好出现一次 。

	任务是找出重复的数字a 和缺失的数字 b 。

	返回一个下标从 0 开始、长度为 2 的整数数组 ans ，其中 ans[0] 等于 a ，ans[1] 等于 b 。
*/

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	cnt := make([]int, n*n+1)
	for _, row := range grid {
		for _, x := range row {
			cnt[x]++
		}
	}

	ans := make([]int, 2)
	for i := 1; i <= n*n; i++ {
		if cnt[i] == 2 {
			ans[0] = i
		} else if cnt[i] == 0 {
			ans[1] = i
		}
	}
	return ans
}
