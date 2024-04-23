package main

/*
	题干：
	有一个书店老板，他的书店开了 n 分钟。每分钟都有一些顾客进入这家商店。给定一个长度为 n 的整数数组 customers ，
	其中 customers[i] 是在第 i 分钟开始时进入商店的顾客数量，所有这些顾客在第 i 分钟结束后离开。

	在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。

	当书店老板生气时，那一分钟的顾客就会不满意，若老板不生气则顾客是满意的。

	书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 minutes 分钟不生气，但却只能使用一次。

	请你返回 这一天营业下来，最多有多少客户能够感到满意 。
*/

/*
	本题可以拆分成两个问题：
	1. 老板不生气时顾客数量之和 s0。这些顾客可以感到满意。
	2. 长度为 minutes 的连续子数组，老板生气时的顾客数量之和 s1 的最大值 maxS1。这些顾客可以感到满意。

	最终答案为 s0 + maxS1

	第二个问题可以用定长滑动窗口解决。
	例如数组 customers = [3,1,4,1,5,9] 和 minutes = 3，假设老板一直在生气：
	1. 计算第一个长为 3 的子数组的元素和 3 + 1 + 4 = 8。
	2. 计算第二个长为 3 的子数组的元素和，我们可以在第一个子数组的基础上，增加 customers[3] = 1，减少 customers[0] = 3，得到 8 + 1 - 3 = 6。
	3. 计算第三个长为 3 的子数组的元素和，我们可以在第二个子数组的基础上，增加 customers[4] = 5，减少 customers[1] = 1，得到 6 + 5 - 1 = 10。
	4. 计算第四个长为 3 的子数组的元素和，我们可以在第三个子数组的基础上，增加 customers[5] = 9，减少 customers[2] = 4，得到 10 + 9 - 4 = 15。

	最大的长为 3 的子数组和为 15。
*/

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	s := [2]int{}
	maxS1 := 0
	for i, c := range customers {
		s[grumpy[i]] += c
		if i < minutes-1 { // 窗口长度不足 minutes
			continue
		}
		maxS1 = max(maxS1, s[1])
		if grumpy[i-minutes+1] > 0 {
			s[1] -= customers[i-minutes+1] // 窗口最左边元素离开窗口
		}
	}
	return s[0] + maxS1
}
