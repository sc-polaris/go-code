package main

import (
	"fmt"
	"io"
)

const (
	N = 110
	M = 11
)

var n, m int

// 方案1：O(n^2*m)
var f [N][M]int // 测量区间为i，j个鸡蛋在最坏情况下测量的最小方案数量

func solve1() {
	for {
		_, err := fmt.Scanf("%d %d", &n, &m)
		if err == io.EOF {
			break
		}

		// 1个鸡蛋测量i层楼的最小方案数
		for i := 1; i <= n; i++ {
			f[i][1] = i
		}
		// i个鸡蛋测量1层楼的最小方案数
		for i := 1; i <= m; i++ {
			f[1][i] = 1
		}
		for i := 2; i <= n; i++ {
			for j := 2; j <= m; j++ {
				f[i][j] = f[i][j-1] // 不用第j个鸡蛋
				// 共有1～i层可以扔鸡蛋，分碎和不碎两种情况
				for k := 1; k <= i; k++ {
					// 蛋碎，搜索区间变成1~k-1，鸡蛋个数减一，方案数为f[k - 1, j - 1]
					// 蛋没碎，搜索区间变成k+1~i，第j个蛋可重复利用，方案数为f[i - k, j]
					// 加上第k层
					f[i][j] = min(f[i][j], max(f[k-1][j-1], f[i-k][j])+1)
				}
			}
		}
		fmt.Println(f[n][m])
	}
}

// 方案2 O(nm)
var dp [N][M]int // j个鸡蛋测量i次最大能测量的区间长度

func solve2() {
	for {
		_, err := fmt.Scanf("%d %d", &n, &m)
		if err == io.EOF {
			break
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				// f[i-1][j-1]：鸡蛋碎了 f[i-1][j]：鸡蛋没碎
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + 1
			}
			if dp[i][m] >= n {
				fmt.Println(i)
				break
			}
		}
	}
}

func main() {
	//solve1()
	solve2()
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
