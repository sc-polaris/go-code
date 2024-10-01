package main

import "math/bits"

/*
	给定一个正整数 n ，请你统计在 [0, n] 范围的非负整数中，有多少个整数的二进制表示中不存在 连续的 1 。

	定义 dfs(i,pre1, isLimit,isNum) 表示构造第 i 位及其之后数位的合法方案数，其余参数的含义为：
	· pre1 表示前一个比特位是否为 1。如果是 1，则当前比特位不能填 1。
	· isLimit 表示当前是否受到了 n 的约束。若为真，则第 i 位填入的数字至多为 n>>i&1，否则可以是 1。设填入的数字至多为 up，
	  如果在 isLimit 为真的情况下填了 up，那么后续填入的数字仍会受到 n 的约束。
	· isNum 表示 i 前面的数位是否填了数字。若为假，则当前位可以跳过（不填数字），或者要填入的数字至少为 1；若为真，则要填入的
	  数字可以从 0 开始。前导零对答案无影响，isNum 可以省略。

	代码实现时，由于（i,pre1,true）这种状态在整个递归过程中至多出现一次，没必要记忆化。或者说，我们只需要记忆化（i,pre1,false)
	这样的状态，也就是在 isLimit=false 时才去记忆化。

	什么情况下必须要有 isNum 参数？
	考虑这样一个问题，计算 [0,n] 中，有多少个数 x 满足：统计 x 的每个数位，要求 0,1,2,...,9 的出现次数都是偶数。
	这里如果把前导零也统计进去的话，就会和 x 中的 0 混在一起了，没法判断 x 中的 0 是否出现了偶数次。
*/

func findIntegers(n int) int {
	m := bits.Len(uint(n))
	memo := make([][2]int, m)
	for i := range memo {
		memo[i] = [2]int{-1, -1} // -1 表示没有计算过
	}

	// pre 表示上一个比特位填的数字
	var dfs func(int, int, bool) int
	dfs = func(i, pre int, isLimit bool) (res int) {
		if i < 0 {
			return 1
		}
		if !isLimit {
			p := &memo[i][pre]
			if *p >= 0 { // 之前计算过
				return *p
			}
			defer func() { *p = res }() // 记忆化
		}
		up := 1
		if isLimit {
			up = n >> i & 1
		}
		res = dfs(i-1, 0, isLimit && up == 0) // 填 0
		if pre == 0 && up == 1 {
			res += dfs(i-1, 1, isLimit) // 填 1
		}
		return
	}
	return dfs(m-1, 0, true) // 从高位到低位
}
