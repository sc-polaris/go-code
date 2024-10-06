package main

/*
	在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

	你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

	给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
*/

func canCompleteCircuit(gas []int, cost []int) int {
	var ans, minS, s int // s 表示油量，minS 表示最小油量
	for i, g := range gas {
		s += g - cost[i] // 在 i 处加油，然后从 i 到 i+1
		if s < minS {
			minS = s    // 更新最小油量
			ans = i + 1 // 注意 s 减去 c 之后，汽车在 i+1 而不是 i
		}
	}
	// 循环结束后，s 即为 gas 之和减去 cost 之和
	if s < 0 {
		return -1
	}
	return ans
}
