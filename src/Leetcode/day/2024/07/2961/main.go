package main

/*
	给你一个下标从 0 开始的二维数组 variables ，其中 variables[i] = [ai, bi, ci, mi]，以及一个整数 target 。

	如果满足以下公式，则下标 i 是 好下标：
	· 0 <= i < variables.length
	· ((ai^bi % 10)^ci) % mi == target
	返回一个由 好下标 组成的数组，顺序不限 。
*/

func getGoodIndices(variables [][]int, target int) (ans []int) {
	for i, v := range variables {
		if pow(pow(v[0], v[1], 10), v[2], v[3]) == target {
			ans = append(ans, i)
		}
	}
	return
}

func pow(x, n, mod int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
