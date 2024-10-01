package main

import "fmt"

/*
	给你一个二维 3 x 3 的矩阵 grid ，每个格子都是一个字符，要么是 'B' ，要么是 'W' 。字符 'W' 表示白色，字符 'B' 表示黑色。

	你的任务是改变 至多一个 格子的颜色，使得矩阵中存在一个 2 x 2 颜色完全相同的正方形。

	如果可以得到一个相同颜色的 2 x 2 正方形，那么返回 true ，否则返回 false 。
*/

/*
	遍历矩阵中的每个 2×2 子矩形。
	对于每个子矩形，统计 B 和 W 的个数，如果其中一个字母的出现次数 ≥3，则返回 true。
	注：也可以判断其中一个字母的出现次数 / 2
	如果四个子矩形都不满足要求，返回 false。
	代码实现时，由于 B 和 W 的 ASCII 值的奇偶性（二进制最低位）不同，可以统计其二进制最低位，代替统计字母。
*/

func canMakeSquare(grid [][]byte) bool {
	check := func(i, j int) bool {
		cnt := [2]int{}
		cnt[grid[i][j]&1]++
		cnt[grid[i][j+1]&1]++
		cnt[grid[i+1][j]&1]++
		cnt[grid[i+1][j+1]&1]++
		return cnt[0] != 2
	}
	return check(0, 0) || check(0, 1) || check(1, 0) || check(1, 1)
}

func canMakeSquare2(grid [][]byte) bool {
	check := func(i, j int) bool {
		count := 0
		for x := 0; x <= 1; x++ {
			for y := 0; y <= 1; y++ {
				if grid[i+x][j+y] == 'B' {
					count++
				}
			}
		}
		fmt.Println(count)
		return count != 2
	}
	return check(0, 0) || check(0, 1) || check(1, 0) || check(1, 1)
}
