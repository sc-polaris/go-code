package main

/*
	现有一份 n + m 次投掷单个 六面 骰子的观测数据，骰子的每个面从 1 到 6 编号。观测数据中缺失了 n 份，你手上只拿到剩余 m
	次投掷的数据。幸好你有之前计算过的这 n + m 次投掷数据的 平均值 。

	给你一个长度为 m 的整数数组 rolls ，其中 rolls[i] 是第 i 次观测的值。同时给你两个整数 mean 和 n 。

	返回一个长度为 n 的数组，包含所有缺失的观测数据，且满足这 n + m 次投掷的 平均值 是 mean 。如果存在多组符合要求的答案，
	只需要返回其中任意一组即可。如果不存在答案，返回一个空数组。

	k 个数字的 平均值 为这些数字求和后再除以 k 。

	注意 mean 是一个整数，所以 n + m 次投掷的总和需要被 n + m 整除。
*/

/*
	n+m 次投掷的总和为 mean * (n+m)
	其中已知数据和为	rem = mean*(n+m)-s

	如果 rem 不在区间 [n,6n] 中，则答案不存在，否则构造方案如下：
	· 计算 avg = ⌊rem/n⌋，extra = rem mod n。
	· 答案中有 extra 个 avg+1 以及 n-extra 个 avg。

	例如 rem = 20，n = 6，计算平均值下取整，的 avg = ⌊20/6⌋ = 3。如果答案为 [3,3,3,3,3,3]，就还少了 2，那么把前两个数加一，
	得 [4,4,3,3,3,3]，即为答案

*/

func missingRolls(rolls []int, mean int, n int) []int {
	rem := mean * (n + len(rolls))
	for _, roll := range rolls {
		rem -= roll
	}
	if rem < n || rem > n*6 {
		return nil
	}

	avg, extra := rem/n, rem%n
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = avg + 1
	}
	for i := extra; i < n; i++ {
		ans[i] = avg
	}
	return ans
}
