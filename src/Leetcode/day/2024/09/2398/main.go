package main

/*
	你有 n 个机器人，给你两个下标从 0 开始的整数数组 chargeTimes 和 runningCosts ，两者长度都为 n 。第 i 个机器人充电
	时间为 chargeTimes[i] 单位时间，花费 runningCosts[i] 单位时间运行。再给你一个整数 budget 。

	运行 k 个机器人 总开销 是 max(chargeTimes) + k * sum(runningCosts) ，其中 max(chargeTimes) 是这 k 个机器人
	中最大充电时间，sum(runningCosts) 是这 k 个机器人的运行时间之和。

	请你返回在 不超过 budget 的前提下，你 最多 可以 连续 运行的机器人数目为多少。
*/

/*
	题目要求机器人连续运行，看成一个连续子数组，题目要求计算最长子数组长度。

	枚举子数组右端点 right，我们需要知道此时左端点 left 的最小值，这样子数组尽量长。

	由于有 budget 的限制，所以 right 越大，left 也越大，有单调性，可以用滑动窗口解决。

	本题的一种做法是二分答案，这样就转换成了固定长度的 239 题。

	但实际上不用二分，在 239 题的基础上，把定长滑窗改为不定长滑窗，套路如下：
	1. 入：chargeTimes[right] 进入窗口时，弹出队尾的 ≤chargeTimes[right] 的元素。
	2. 出：如果总开销超过 budget，则不断移出左端点，直到总开销不超过 budget。特别地，如果左端点恰好等于队首，则弹出队首。
	3. 更新答案：用窗口长度 right−left+1 更新答案的最大值。
	⚠注意：为了方便判断队首是否要出队，单调队列中保存的是下标。
*/

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) (ans int) {
	var q []int
	sum := int64(0)
	l := 0
	for r, t := range chargeTimes {
		// 1. 入
		for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		sum += int64(runningCosts[r])

		// 2. 出
		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(r-l+1)*sum > budget {
			if q[0] == l {
				q = q[1:]
			}
			sum -= int64(runningCosts[l])
			l++
		}

		// 3. 更新答案
		ans = max(ans, r-l+1)
	}
	return
}
