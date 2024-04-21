package main

import "slices"

/*
	枚举选哪个
	剪枝：
	最大 d 个数的和
	i + i-1 + i-2 + ... + i-d+1 = (i+i-d+1)*d/2 = (i*2-d+1)*d/2
*/

func combinationSum3(k int, n int) (ans [][]int) {
	var path []int
	var dfs func(int, int)
	dfs = func(i, t int) {
		d := k - len(path)              // 还需要选 d 个数
		if t < 0 || t > (i*2-d+1)*d/2 { // 剪枝
			return
		}
		if d == 0 { // 找到一个合法组合
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := i; j >= d; j-- {
			path = append(path, j)
			dfs(j-1, t-j)
			path = path[:len(path)-1] // 回溯
		}
	}
	dfs(9, n)
	return
}
