package main

import "fmt"

/*
	给你一个正整数 n。

	如果一个二进制字符串 x 的所有长度为 2 的子字符串中包含 至少 一个 "1"，则称 x 是一个 有效 字符串。

	返回所有长度为 n 的 有效 字符串，可以以任意顺序排列。
*/

// 回溯
func validStrings(n int) (ans []string) {
	path := make([]byte, n)
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path)) // 注意 string(path) 需要 O(n) 时间
			return
		}

		path[i] = '1'
		dfs(i + 1)
		if i == 0 || path[i-1] == '1' {
			path[i] = '0' // 直接覆盖
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}

/*
	位运算
	怎么判断二进制中是否有相邻的 0？
	我们可以把 i 取反（保留低 n 位），记作 x。问题变成：判断 x 中是否有相邻的 1。
	需要一个一个地遍历二进制数 x 的每一位吗？
	不需要，我们可以用 x & (x >> 1) 来判断，如果这个值不为零，则说明 x 中有相邻的 1，反之没有。例如 x=110，右移一位得 011，
	可以发现这两个二进制数的次低位都是 1，所以计算 AND 的结果必然不为 0。

	代码实现时，可以直接枚举取反后的值 x，如果 x & (x >> 1) 等于 0，就把 x 取反后的值（也就是 i）加入答案。
	如何取反？
	1. 创建一个低 n 位全为 1 的二进制数 mask = (1 << n) - 1。
	2. 计算 x ^ mask，由于 0 和 1 异或后变成了 1，1 和 1 异或后变成了 0，所以 x 的低 n 位都取反了。


*/

func validStrings2(n int) (ans []string) {
	mask := 1<<n - 1
	for x := range 1 << n {
		if x>>1&x == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, x^mask))
		}
	}
	return
}
